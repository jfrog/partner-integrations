const { blocks, mktext } = require('./common');

module.exports = (groupings, jpdOrigin, policyOrWatchName, forcedSummaryFormat = false) => [{
  blocks: blocks({
    count: 5,
    groupings,
    jpdOrigin,
    fieldFunction: (issue) => ({
      Issue: `<https://nvd.nist.gov/vuln/detail/${issue.id}|${issue.id}>`,
      File: issue.file,
    }),
  }, policyOrWatchName, forcedSummaryFormat),
  text: mktext({
    emoji: ':rage:',
    groupings,
  }),
}];
