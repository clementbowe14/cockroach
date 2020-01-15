// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package cli

import (
	"context"
	"os"
	"sort"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/testutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/stretchr/testify/assert"
)

// TestZipContainsAllInternalTables verifies that we don't add new internal tables
// without also taking them into account in a `debug zip`. If this test fails,
// add your table to either of the []string slices referenced in the test (which
// are used by `debug zip`) or add it as an exception after having verified that
// it indeed should not be collected (this is rare).
// NB: if you're adding a new one, you'll also have to update TestZip.
func TestZipContainsAllInternalTables(t *testing.T) {
	defer leaktest.AfterTest(t)()

	s, db, _ := serverutils.StartServer(t, base.TestServerArgs{})
	defer s.Stopper().Stop(context.Background())

	rows, err := db.Query(`
SELECT concat('crdb_internal.', table_name) as name FROM [ SHOW TABLES FROM crdb_internal ] WHERE
    table_name NOT IN (
-- whitelisted tables that don't need to be in debug zip
'backward_dependencies',
'builtin_functions',
'create_statements',
'forward_dependencies',
'index_columns',
'table_columns',
'table_indexes',
'ranges',
'ranges_no_leases',
'predefined_comments',
'session_trace',
'session_variables',
'tables'
)
ORDER BY name ASC`)
	assert.NoError(t, err)

	var tables []string
	for rows.Next() {
		var table string
		assert.NoError(t, rows.Scan(&table))
		tables = append(tables, table)
	}
	tables = append(tables, "system.jobs", "system.descriptor", "system.namespace")
	sort.Strings(tables)

	var exp []string
	exp = append(exp, debugZipTablesPerNode...)
	exp = append(exp, debugZipTablesPerCluster...)
	sort.Strings(exp)

	assert.Equal(t, exp, tables)
}

func TestZip(t *testing.T) {
	defer leaktest.AfterTest(t)()

	dir, cleanupFn := testutils.TempDir(t)
	defer cleanupFn()

	c := newCLITest(cliTestParams{
		storeSpecs: []base.StoreSpec{{
			Path: dir,
		}},
	})
	defer c.cleanup()

	out, err := c.RunWithCapture("debug zip " + os.DevNull)
	if err != nil {
		t.Fatal(err)
	}

	const expected = `debug zip ` + os.DevNull + `
writing ` + os.DevNull + `
  debug/events.json
  debug/rangelog.json
  debug/liveness.json
  debug/settings.json
  debug/reports/problemranges.json
  debug/crdb_internal.cluster_queries.txt
  debug/crdb_internal.cluster_sessions.txt
  debug/crdb_internal.cluster_settings.txt
  debug/crdb_internal.jobs.txt
  debug/system.jobs.txt
  debug/system.descriptor.txt
  debug/system.namespace.txt
  debug/crdb_internal.kv_node_status.txt
  debug/crdb_internal.kv_store_status.txt
  debug/crdb_internal.schema_changes.txt
  debug/crdb_internal.partitions.txt
  debug/crdb_internal.zones.txt
  debug/nodes/1/status.json
  debug/nodes/1/crdb_internal.feature_usage.txt
  debug/nodes/1/crdb_internal.gossip_alerts.txt
  debug/nodes/1/crdb_internal.gossip_liveness.txt
  debug/nodes/1/crdb_internal.gossip_network.txt
  debug/nodes/1/crdb_internal.gossip_nodes.txt
  debug/nodes/1/crdb_internal.leases.txt
  debug/nodes/1/crdb_internal.node_build_info.txt
  debug/nodes/1/crdb_internal.node_metrics.txt
  debug/nodes/1/crdb_internal.node_queries.txt
  debug/nodes/1/crdb_internal.node_runtime_info.txt
  debug/nodes/1/crdb_internal.node_sessions.txt
  debug/nodes/1/crdb_internal.node_statement_statistics.txt
  debug/nodes/1/crdb_internal.node_txn_stats.txt
  debug/nodes/1/details.json
  debug/nodes/1/gossip.json
  debug/nodes/1/enginestats.json
  debug/nodes/1/stacks.txt
  debug/nodes/1/heap.pprof
  debug/nodes/1/ranges/1.json
  debug/nodes/1/ranges/2.json
  debug/nodes/1/ranges/3.json
  debug/nodes/1/ranges/4.json
  debug/nodes/1/ranges/5.json
  debug/nodes/1/ranges/6.json
  debug/nodes/1/ranges/7.json
  debug/nodes/1/ranges/8.json
  debug/nodes/1/ranges/9.json
  debug/nodes/1/ranges/10.json
  debug/nodes/1/ranges/11.json
  debug/nodes/1/ranges/12.json
  debug/nodes/1/ranges/13.json
  debug/nodes/1/ranges/14.json
  debug/nodes/1/ranges/15.json
  debug/nodes/1/ranges/16.json
  debug/nodes/1/ranges/17.json
  debug/nodes/1/ranges/18.json
  debug/nodes/1/ranges/19.json
  debug/nodes/1/ranges/20.json
  debug/nodes/1/ranges/21.json
  debug/nodes/1/ranges/22.json
  debug/nodes/1/ranges/23.json
  debug/nodes/1/ranges/24.json
  debug/schema/defaultdb@details.json
  debug/schema/postgres@details.json
  debug/schema/system@details.json
  debug/schema/system/comments.json
  debug/schema/system/descriptor.json
  debug/schema/system/eventlog.json
  debug/schema/system/jobs.json
  debug/schema/system/lease.json
  debug/schema/system/locations.json
  debug/schema/system/namespace.json
  debug/schema/system/rangelog.json
  debug/schema/system/replication_constraint_stats.json
  debug/schema/system/replication_critical_localities.json
  debug/schema/system/replication_stats.json
  debug/schema/system/reports_meta.json
  debug/schema/system/role_members.json
  debug/schema/system/settings.json
  debug/schema/system/table_statistics.json
  debug/schema/system/ui.json
  debug/schema/system/users.json
  debug/schema/system/web_sessions.json
  debug/schema/system/zones.json
`

	assert.Equal(t, expected, out)
}
