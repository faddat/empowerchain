/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    "node_modules/flowbite-vue/**/*.{js,jsx,ts,tsx}",
    "node_modules/flowbite/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        lightBlack: "rgba(0, 0, 0, 0.22)",
        black: "rgba(0, 0, 0, 1)",
        lightGreen: "rgba(0, 211, 73, 1)",
        lightGray: "rgba(255, 255, 255, 0.2)",
        gray: "rgba(112, 112, 112, 1)",
        white: "rgba(255, 255, 255, 1)",
        lightWhite: "rgba(255, 255, 255, 0.08)",
      },
      fontFamily: {
        openSans: ["Open Sans", "sans-serif"],
      },
      backgroundImage: {
        "main-image": "url('./src/assets/images/main-bg.jpg')",
      },
    },
  },
  plugins: [require("flowbite/plugin")],
};
