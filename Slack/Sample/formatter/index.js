const { lookupFormatMapper } = require('./shared');

const SEVERITY_MAPPING = {
  High: 0,
  Medium: 1,
  Low: 2,
  Unknown: 3,
};

const SLACK_APP_MAX_ISSUES_PER_ENTRY = 40;
const SUMMARY_MESSAGE_FORMAT = 'summary';
const ISSUE_MESSAGE_FORMAT = 'issue';

const normalize = (violation) => violation.issues
  .map((issue) => issue.impacted_artifacts.map((artifact) => artifact.infected_files.map((file) => ({
    watch_name: violation.watch_name,
    severity: issue.severity,
    type: issue.type,
    pkg_type: artifact.pkg_type,
    summary: issue.summary,
    path: `${artifact.path.replace('default/', '')}`,
    file: file.name || artifact.name,
    description: issue.description,
    id: issue.cve || issue.summary,
  })))).flat(4);

module.exports = ({ violation, jpdOrigin }, format, policyOrWatchName) => {
  const normalizedViolations = normalize(violation);
  const reports = normalizedViolations.sort((a, b) => SEVERITY_MAPPING[a.severity] - SEVERITY_MAPPING[b.severity]);

  let messageFormat = [SUMMARY_MESSAGE_FORMAT, ISSUE_MESSAGE_FORMAT].includes(format) ? `${format}` : SUMMARY_MESSAGE_FORMAT;
  let forcedSummaryFormat = false;

  // if there're too many violations then use 'summary' format
  if (messageFormat === ISSUE_MESSAGE_FORMAT && reports.length > SLACK_APP_MAX_ISSUES_PER_ENTRY) {
    console.log(`Number of violations in a report exceeded max allowed for individual message. Switching to '${SUMMARY_MESSAGE_FORMAT}' format.`, { reportsLength: reports.length, SLACK_APP_MAX_ISSUES_PER_ENTRY });
    messageFormat = SUMMARY_MESSAGE_FORMAT;
    forcedSummaryFormat = true;
  }

  const mapper = lookupFormatMapper(messageFormat);
  return mapper(reports, jpdOrigin, policyOrWatchName, forcedSummaryFormat)?.map((r) => ({
    format,
    ...r,
  }));
};
