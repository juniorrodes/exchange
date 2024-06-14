/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      "./components/*.{templ,js,go,html}"
  ],
  theme: {
    container: {
        center: true,
    },
    extend: {},
  },
  plugins: [
      require('@tailwindcss/forms'),
  ],
}

