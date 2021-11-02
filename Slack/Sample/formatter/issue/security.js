const { blocks, fields, longFields } = require('./common');

module.exports = (violations, jpdOrigin, policyOrWatchName) => violations.map((violation) => blocks({
  report: violation,
  jpdOrigin,
  emoji: ':rage:',
  fields: fields({
    Severity: violation.severity,
    Type: violation.pkg_type,
    Issue: `<https://nvd.nist.gov/vuln/detail/${violation.id}|${violation.id}>`,
    File: violation.file,
  }),
  longFields: longFields({
    Summary: `\n${violation.summary.replace(violation.id, '')}`,
    Path: `\n${violation.path}`,
  }),
}, policyOrWatchName));
