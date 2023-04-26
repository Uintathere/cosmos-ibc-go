// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require("prism-react-renderer/themes/github");
const darkCodeTheme = require("prism-react-renderer/themes/dracula");

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "IBC-Go Documentation",
  tagline: "Documentation for IBC-Go",
  favicon: "img/white-cosmos-icon.svg",

  // Set the production url of your site here
  url: "https://ibc.cosmos.network/",
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: "/",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "cosmos", // Usually your GitHub org/user name.
  projectName: "ibc-go", // Usually your repo name.

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve("./sidebars.js"),
          // Routed the docs to the root path
          routeBasePath: "/",
          // Exclude template markdown files from the docs
          exclude: ["**/*.template.md"],
        },
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      image: "img/ibc-go-image.png",
      navbar: {
        title: "IBC-Go",
        logo: {
          alt: "Black Cosmos Icon",
          src: "img/ibc.svg",
        },
        items: [
          {
            type: "docSidebar",
            sidebarId: "defaultSidebar",
            position: "left",
            label: "Documentation",
          },
          {
            type: "doc",
            position: "left",
            docId: "README",
            docsPluginId: "adrs",
            label: "Architecture Decision Records",
          },
          {
            type: "docsVersionDropdown",
            position: "right",
            dropdownActiveClassDisabled: true,
            // TODO: versions not yet migrated to docusaurus
            dropdownItemsAfter: [
              {
                href: "https://ibc.cosmos.network/v7.0.0/",
                label: "pre v7.1",
                target: "_self",
              },
            ],
          },
          {
            href: "https://github.com/cosmos/ibc-go",
            html: `<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="github-icon">
            <path fill-rule="evenodd" clip-rule="evenodd" d="M12 0.300049C5.4 0.300049 0 5.70005 0 12.3001C0 17.6001 3.4 22.1001 8.2 23.7001C8.8 23.8001 9 23.4001 9 23.1001C9 22.8001 9 22.1001 9 21.1001C5.7 21.8001 5 19.5001 5 19.5001C4.5 18.1001 3.7 17.7001 3.7 17.7001C2.5 17.0001 3.7 17.0001 3.7 17.0001C4.9 17.1001 5.5 18.2001 5.5 18.2001C6.6 20.0001 8.3 19.5001 9 19.2001C9.1 18.4001 9.4 17.9001 9.8 17.6001C7.1 17.3001 4.3 16.3001 4.3 11.7001C4.3 10.4001 4.8 9.30005 5.5 8.50005C5.5 8.10005 5 6.90005 5.7 5.30005C5.7 5.30005 6.7 5.00005 9 6.50005C10 6.20005 11 6.10005 12 6.10005C13 6.10005 14 6.20005 15 6.50005C17.3 4.90005 18.3 5.30005 18.3 5.30005C19 7.00005 18.5 8.20005 18.4 8.50005C19.2 9.30005 19.6 10.4001 19.6 11.7001C19.6 16.3001 16.8 17.3001 14.1 17.6001C14.5 18.0001 14.9 18.7001 14.9 19.8001C14.9 21.4001 14.9 22.7001 14.9 23.1001C14.9 23.4001 15.1 23.8001 15.7 23.7001C20.5 22.1001 23.9 17.6001 23.9 12.3001C24 5.70005 18.6 0.300049 12 0.300049Z" fill="currentColor"/>
            </svg>
            `,
            position: "right",
          },
        ],
      },
      footer: {
        style: "dark",
        links: [
          {
            title: "Documentation",
            items: [
              {
                label: "Cosmos Hub",
                href: "https://hub.cosmos.network",
              },
              {
                label: "CometBFT",
                href: "https://docs.cometbft.com",
              },
              {
                label: "IBC Go",
                href: "https://ibc.cosmos.network",
              },
            ],
          },
          {
            title: "Community",
            items: [
              {
                label: "Discord",
                href: "https://discord.gg/Wtmk6ZNa8G",
              },
              {
                label: "Twitter",
                href: "https://twitter.com/cosmosibc",
              },
            ],
          },
          {
            title: "More",
            items: [
              {
                label: "GitHub",
                href: "https://github.com/cosmos/ibc-go",
              },
            ],
          },
        ],
        copyright: `<p>The development of IBC-Go is led primarily by <a href="https://interchain.berlin/">Interchain GmbH</a>. Funding for this development comes primarily from the Interchain Foundation, a Swiss non-profit.</p>`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
  themes: ["@you54f/theme-github-codeblock"],
  plugins: [
    [
      "@docusaurus/plugin-content-docs",
      {
        id: "adrs",
        path: "architecture",
        routeBasePath: "architecture",
        sidebarPath: require.resolve("./sidebars.js"),
        exclude: ["**/*.template.md"],
      },
    ],
    async function myPlugin(context, options) {
      return {
        name: "docusaurus-tailwindcss",
        configurePostCss(postcssOptions) {
          postcssOptions.plugins.push(require("postcss-import"));
          postcssOptions.plugins.push(require("tailwindcss/nesting"));
          postcssOptions.plugins.push(require("tailwindcss"));
          postcssOptions.plugins.push(require("autoprefixer"));
          return postcssOptions;
        },
      };
    },
  ],
};

module.exports = config;
