package yandex

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	lt "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1"
	ltagent "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/agent"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

const loadtestingAgentResource = "yandex_loadtesting_agent.test-lt-agent"
const loadtestingAgentSubnetResource = "yandex_vpc_subnet.loadtesting-agent-test-subnet"

func init() {
	resource.AddTestSweepers("yandex_loadtesting_agent", &resource.Sweeper{
		Name: "yandex_loadtesting_agent",
		F:    testSweepLoadtestingAgents,
		Dependencies: []string{
			"yandex_vpc_network",
			"yandex_vpc_subnet",
			"yandex_iam_service_account",
			"yandex_compute_instance",
		},
	})
}

func testSweepLoadtestingAgents(_ string) error {
	conf, err := configForSweepers()
	if err != nil {
		return fmt.Errorf("error getting client: %s", err)
	}

	result := &multierror.Error{}
	req := &lt.ListAgentsRequest{FolderId: conf.FolderID}
	it := conf.sdk.Loadtesting().Agent().AgentIterator(conf.Context(), req)
	for it.Next() {
		id := it.Value().GetId()
		if !sweepLoadtestingAgent(conf, id) {
			result = multierror.Append(result, fmt.Errorf("failed to sweep Loadtesting Agent %q", id))
		}
	}

	return result.ErrorOrNil()
}

func sweepLoadtestingAgent(conf *Config, id string) bool {
	return sweepWithRetry(sweepLoadtestingAgentOnce, conf, "Loadtesting Agent", id)
}

func sweepLoadtestingAgentOnce(conf *Config, id string) error {
	ctx, cancel := conf.ContextWithTimeout(yandexIAMServiceAccountDefaultTimeout)
	defer cancel()

	op, err := conf.sdk.Loadtesting().Agent().Delete(ctx, &lt.DeleteAgentRequest{
		AgentId: id,
	})
	return handleSweepOperation(ctx, conf, op, err)
}

func TestAccResourceLoadtestingAgent_full(t *testing.T) {
	t.Parallel()

	var agent ltagent.Agent
	var instance compute.Instance
	agentName := acctest.RandomWithPrefix("tf-loadtesting-agent")
	agentDescription := "Description for test full"
	folderID := getExampleFolderID()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLoadtestingAgentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLoadtestingAgentFull(agentName, agentDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLoadtestingAgentExists(loadtestingAgentResource, &agent),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "name", agentName),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "description", agentDescription),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "folder_id", folderID),
					resource.TestCheckResourceAttrSet(loadtestingAgentResource, "compute_instance_id"),
					testAccCheckLoadtestingComputeInstanceExists(&agent, &instance),
					resource.TestCheckResourceAttrSet(loadtestingAgentResource, "compute_instance.0.service_account_id"),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "compute_instance.0.resources.0.memory", "4"),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "compute_instance.0.resources.0.cores", "4"),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "compute_instance.0.labels.purpose", "loadtesting-agent"),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "compute_instance.0.metadata.field1", "metavalue1"),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "compute_instance.0.metadata.field2", "other value 2"),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "compute_instance.0.computed_metadata.field1", "metavalue1"),
					resource.TestCheckResourceAttr(loadtestingAgentResource, "compute_instance.0.computed_metadata.field2", "other value 2"),
					resource.TestCheckResourceAttrSet(loadtestingAgentResource, "compute_instance.0.computed_metadata.loadtesting-created"),
					testAccCheckLoadtestingComputeInstanceHasResources(&instance, 4, 50, 4.0),
					testAccCheckLoadtestingComputeInstanceLabel(&instance, "purpose", "loadtesting-agent"),
					testAccCheckLoadtestingComputeInstanceMetadata(&instance, "field1", "metavalue1"),
					testAccCheckLoadtestingComputeInstanceMetadata(&instance, "field2", "other value 2"),
					testAccCheckLoadtestingComputeBootDiskExists(&instance, true, 17),
					testAccCheckLoadtestingComputeNetworkInterface(loadtestingAgentSubnetResource, &instance, true, false),
				),
			},
			{
				ResourceName:            loadtestingAgentResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"compute_instance.0.metadata"},
			},
		},
	})
}

func testAccCheckLoadtestingAgentDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "yandex_loadtesting_agent" {
			continue
		}

		_, err := config.sdk.Loadtesting().Agent().Get(context.Background(), &lt.GetAgentRequest{
			AgentId: rs.Primary.ID,
		})
		if err == nil {
			return fmt.Errorf("Loadtesting Agent still exists")
		}
	}

	return nil
}

func testAccCheckLoadtestingAgentExists(n string, agent *ltagent.Agent) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)

		found, err := config.sdk.Loadtesting().Agent().Get(context.Background(), &lt.GetAgentRequest{
			AgentId: rs.Primary.ID,
		})
		if err != nil {
			return err
		}

		if found.Id != rs.Primary.ID {
			return fmt.Errorf("Loadtesting Agent not found")
		}

		*agent = *found

		return nil
	}
}

func testAccCheckLoadtestingComputeInstanceExists(agent *ltagent.Agent, instance *compute.Instance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		found, err := config.sdk.Compute().Instance().Get(context.Background(), &compute.GetInstanceRequest{
			InstanceId: agent.GetComputeInstanceId(),
			View:       compute.InstanceView_FULL,
		})
		if err != nil {
			return err
		}

		if found.Id != agent.GetComputeInstanceId() {
			return fmt.Errorf("Compute instance not found")
		}

		*instance = *found

		return nil
	}
}

func testAccCheckLoadtestingComputeBootDiskExists(instance *compute.Instance, autodelete bool, size int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		found, err := config.sdk.Compute().Disk().Get(context.Background(), &compute.GetDiskRequest{
			DiskId: instance.GetBootDisk().GetDiskId(),
		})
		if err != nil {
			return err
		}

		if found.Id != instance.GetBootDisk().GetDiskId() {
			return fmt.Errorf("Compute Instance Disk not found")
		}

		if instance.GetBootDisk().AutoDelete != autodelete {
			return fmt.Errorf("Wrong instance boot disk autodelete: expected %v, got %v", autodelete, instance.GetBootDisk().AutoDelete)
		}
		if found.Size != toBytes(size) {
			return fmt.Errorf("Wrong instance boot disk size: expected %d, got %d", toBytes(size), found.Size)
		}

		return nil
	}
}

func testAccCheckLoadtestingComputeInstanceHasResources(instance *compute.Instance, cores, coreFraction int64, memoryGB float64) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resources := instance.GetResources()
		if resources.Cores != cores {
			return fmt.Errorf("Wrong instance Cores resource: expected %d, got %d", cores, resources.Cores)
		}
		if resources.CoreFraction != coreFraction {
			return fmt.Errorf("Wrong instance Cores Fraction resource: expected %d, got %d", coreFraction, resources.CoreFraction)
		}
		memoryBytes := toBytesFromFloat(memoryGB)
		if resources.Memory != memoryBytes {
			return fmt.Errorf("Wrong instance Memory resource: expected %f, got %d", memoryGB, toGigabytes(resources.Memory))
		}
		return nil
	}
}

func testAccCheckLoadtestingComputeInstanceLabel(instance *compute.Instance, key string, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if instance.Labels == nil {
			return fmt.Errorf("no labels found on instance %s", instance.Name)
		}

		v, ok := instance.Labels[key]
		if !ok {
			return fmt.Errorf("No label found with key %s on instance %s", key, instance.Name)
		}
		if v != value {
			return fmt.Errorf("Expected value '%s' but found value '%s' for label '%s' on instance %s", value, v, key, instance.Name)
		}

		return nil
	}
}

func testAccCheckLoadtestingComputeInstanceMetadata(
	instance *compute.Instance,
	k string, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if instance.Metadata == nil {
			return fmt.Errorf("no metadata")
		}

		mv, ok := instance.Metadata[k]
		if !ok {
			return fmt.Errorf("metadata not found for key '%s'", k)
		}

		if v != mv {
			return fmt.Errorf("bad value for %s: %s", k, mv)
		}

		return nil
	}
}

func testAccCheckLoadtestingComputeNetworkInterface(subnetResourceName string, instance *compute.Instance, hasIPv4Address bool, hasIPv6Address bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[subnetResourceName]
		if !ok {
			return fmt.Errorf("Subnet not found: %s", subnetResourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Subnet ID is not set")
		}

		config := testAccProvider.Meta().(*Config)

		subnet, err := config.sdk.VPC().Subnet().Get(context.Background(), &vpc.GetSubnetRequest{
			SubnetId: rs.Primary.ID,
		})
		if err != nil {
			return err
		}

		if subnet.Id != rs.Primary.ID {
			return fmt.Errorf("Subnet not found")
		}

		if len(instance.NetworkInterfaces) != 1 {
			return fmt.Errorf("Wrong instance network interface count: expected 1, got 0")
		}

		actualHasV4 := instance.NetworkInterfaces[0].PrimaryV4Address != nil
		if actualHasV4 != hasIPv4Address {
			return fmt.Errorf("Wrong instance network interface ipv4: expected %v, got %v", hasIPv4Address, actualHasV4)
		}
		actualHasV6 := instance.NetworkInterfaces[0].PrimaryV6Address != nil
		if actualHasV6 != hasIPv6Address {
			return fmt.Errorf("Wrong instance network interface ipv6: expected %v, got %v", hasIPv6Address, actualHasV6)
		}

		return nil
	}
}

func testAccLoadtestingAgentPrerequisites(name string) string {
	return fmt.Sprintf(`
resource "yandex_vpc_network" "loadtesting-agent-test-network" {}

resource "yandex_vpc_subnet" "loadtesting-agent-test-subnet" {
	zone           = "ru-central1-b"
	network_id     = "${yandex_vpc_network.loadtesting-agent-test-network.id}"
	v4_cidr_blocks = ["192.168.0.0/24"]
}

resource "yandex_iam_service_account" "loadtesting-agent-test-sa" {
	name          = "%s-sa"
}
`, name)
}

func testAccLoadtestingAgentFull(name, desc string) string {
	prereq := testAccLoadtestingAgentPrerequisites(name)
	return fmt.Sprintf(`%s

resource "yandex_loadtesting_agent" "test-lt-agent" {
	name		  = "%s"
	description   = "%s"
		
	compute_instance {
		zone_id = "ru-central1-b"
		service_account_id = "${yandex_iam_service_account.loadtesting-agent-test-sa.id}"
		resources {
			memory = 4
			cores = 4
			core_fraction = 50
		}
		metadata = {
			field1 = "metavalue1"
			field2 = "other value 2"
		}
		labels = {
			purpose = "loadtesting-agent"
		}
		boot_disk {
			initialize_params {
				size = 17
				name = "%s-disk"
				description = "%s-disk-desc"
				block_size = 4096
				type = "network-hdd"
			}
			device_name = "somename"
			auto_delete = true
		}
		network_interface {
			subnet_id = "${yandex_vpc_subnet.loadtesting-agent-test-subnet.id}"
			ipv4 = true
			ipv6 = false
			nat = true  
		}
	}
}
`, prereq, name, desc, name, name)
}
