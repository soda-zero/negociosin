const config = {
  content: [
    "./src/**/*.{html,js,svelte,ts}",
    // 2. Append the path for the Skeleton NPM package and files:
    require("path").join(
      require.resolve("@skeletonlabs/skeleton"),
      "../**/*.{html,js,svelte,ts}"
    ),
  ],
  darkMode: "media",
  theme: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/forms"),
    ...require("@skeletonlabs/skeleton/tailwind/skeleton.cjs")(),
  ],
};

module.exports = config;
