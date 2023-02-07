# Nuxt 3
## Packages
### ESLint
- [Nuxt3 + TypeScript環境でESLint導入](https://qiita.com/shott/items/caa62324ca9827526ba9)
- [Nuxt 3 with ESLint](https://qiita.com/kentarou_masuda/items/c0180fe383b01ba54cbf)

@frontend dir
$ npm install --save-dev eslint

$ npm install --save-dev eslint-plugin-vue

$ npm install --save-dev @typescript-eslint/eslint-plugin @typescript-eslint/parser @nuxtjs/eslint-config-typescript

$ npx eslint --init

✔ How would you like to use ESLint? · problems
✔ What type of modules does your project use? · esm
✔ Which framework does your project use? · vue
✔ Does your project use TypeScript? · Yes
✔ Where does your code run? · browser, node
✔ What format do you want your config file to be in? · JavaScript

$ npx lint