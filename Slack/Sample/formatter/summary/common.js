const _ = require('lodash');
const moment = require('moment');

const shared = require('../shared');

const headerText = (groupings) => `(${Object.entries(groupings).map(([severity, byArt]) => {
  const count = Object.values(byArt).map((it) => it.length).reduce((a, b) => a + b);
  return `${count} ${severity}`;
}).join(', ')})`;

const header = (groupings) => {
  const { type } = (Object.values(groupings).map(Object.values)[0][0][0]);
  const head = headerText(groupings);
  return {
    type: 'section',
    text: {
      type: 'mrkdwn',
      text: `*${_.capitalize(type)} Violation* ${head}`,
    },
  };
};

const timestamp = () => ({
  type: 'context',
  elements: [
    {
      type: 'plain_text',
      text: `Time: ${moment().format('lll')}`,
    },
  ],
});

const getNotificationNameSection = (policyOrWatchName) => ({
  type: 'section',
  text: {
    type: 'mrkdwn',
    text: `*Watch*: ${policyOrWatchName}`,
    verbatim: true,
  },
});

const forcedSummaryFormatMessage = (forcedSummaryFormat) => (forcedSummaryFormat ? {
  type: 'context',
  elements: [
    {
      type: 'mrkdwn',
      text: ':information_source: Too many violations to display in Slack. Use \'Open in platform\' button to see full list of violations.',
    },
  ],
} : []);

const openButton = (originHost, report) => ({
  type: 'actions',
  elements: [{
    type: 'button',
    action_id: 'open-in-platform',
    text: {
      type: 'plain_text',
      text: 'Open in platform',
    },
    url: `${originHost}/ui/watchesNew/edit/${report.watch_name}?activeTab=violations`,
  }],
});

const mktext = ({ emoji, groupings }) => {
  const { type } = (Object.values(groupings).map(Object.values)[0][0][0]);
  return `${emoji} ${_.capitalize(type)} violation: ${headerText(groupings)}`;
};

const summary = (count, severity, violations) => ({
  type: 'section',
  text: {
    type: 'mrkdwn',
    text: `*${severity}* (${Math.min(count, violations.length)} of ${violations.length} shown)`,
  },
});

const artBlocks = (origin, { pkg_type: pkgType, path }) => ({
  type: 'section',
  text: {
    type: 'mrkdwn',
    text: `*${pkgType}*: <${origin}/ui/repos/tree/General/${path}|${path}>`,
  },
});

const formatIssue = (fieldFunction, violations) => ({
  type: 'section',
  fields: violations.map((violation) => shared.fields(fieldFunction(violation))).flat(),
});

const blocks = ({ count, groupings, jpdOrigin, fieldFunction }, policyOrWatchName, forcedSummaryFormat = false) => {
  const head = header(groupings);
  const body = Object.entries(groupings).map(([severity, byArtifact]) => Object.entries(byArtifact).map(([_artifact, violations]) => {
    const sum = summary(count, severity, violations);
    const topFive = violations.slice(0, count);
    const artifactDetails = artBlocks(jpdOrigin, violations[0]);
    const violationDetails = formatIssue(fieldFunction, topFive);
    return [sum].concat(artifactDetails).concat(violationDetails);
  })).flat(2);
  const first = Object.values(Object.values(groupings)[0])[0][0];
  const notificationNameSection = getNotificationNameSection(policyOrWatchName);
  return [head]
    .concat(timestamp())
    .concat(body)
    .concat([notificationNameSection])
    .concat(forcedSummaryFormatMessage(forcedSummaryFormat))
    .concat(openButton(jpdOrigin, first));
};

module.exports = {
  blocks,
  mktext,
  ...require('../shared'),
};
