const _ = require('lodash');

const { lookupType } = require('./common');

module.exports = (reports, jpdOrigin, policyOrWatchName, forcedSummaryFormat = false) => {
  const mapper = lookupType(__dirname, reports[0].type);
  if (mapper) {
    const groupings = Object.entries(_.groupBy(reports, (v) => v.severity))
      .map(([severity, violations]) => ({ [severity]: _.groupBy(violations, (v) => v.path) }))
      .reduce((a, b) => ({ ...a, ...b }));

    const results = mapper(groupings, jpdOrigin, policyOrWatchName, forcedSummaryFormat).map((o) => ({
      ...o,
      type: reports[0].type,
      forcedSummaryFormat,
      issueCount: reports.length,
    }));

    return results;
  }
};
