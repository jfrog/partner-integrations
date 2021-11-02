const { blocks, fields, longFields } = require('./common');

module.exports = (reports, jpdOrigin, policyOrWatchName) => reports.map((report) => blocks({
  report,
  jpdOrigin,
  emoji: ':face_palm:',
  fields: fields({
    Severity: report.severity,
    Type: report.pkg_type,
    'License Name': report.summary,
    File: report.file,
  }),
  longFields: longFields({
    Description: report.description,
    Path: `\n${report.path}`,
  }),
}, policyOrWatchName));
