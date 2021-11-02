const { blocks, mktext } = require('./common');

module.exports = (groupings, jpdOrigin, policyOrWatchName, forcedSummaryFormat = false) => [{
  blocks: blocks({
    count: 5,
    groupings,
    jpdOrigin,
    fieldFunction: (issue) => ({
      License: issue.id,
      File: issue.file,
    }),
  }, policyOrWatchName, forcedSummaryFormat),
  text: mktext({
    emoji: ':face_palm:',
    groupings,
  }),
}];
