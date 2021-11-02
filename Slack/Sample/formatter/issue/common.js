const _ = require('lodash');
const moment = require('moment');

const blocks = ({
  longFields, report, jpdOrigin, fields, emoji,
  header = `${_.capitalize(report.type)} Violation`,
}, policyOrWatchName) => ({
  blocks: [
    {
      type: 'header',
      text: {
        type: 'plain_text',
        text: header,
        emoji: true,
      },
    },
    {
      type: 'divider',
    },
    {
      type: 'section',
      fields,
    }]
    .concat(longFields)
    .concat({
      type: 'section',
      text: {
        type: 'mrkdwn',
        text: `*Watch*: ${policyOrWatchName}`,
        verbatim: true,
      },
    })
    .concat({
      type: 'actions',
      elements: [
        {
          type: 'button',
          action_id: 'open-in-platform',
          text: {
            type: 'plain_text',
            text: 'Open in platform',
          },
          url: `${jpdOrigin}/ui/watchesNew/edit/${report.watch_name}?activeTab=violations`,
        },
      ],
    }),
  text: `${emoji} ${header} (${report.severity})`,
});

module.exports = {
  blocks,
  ...require('../shared'),
};
