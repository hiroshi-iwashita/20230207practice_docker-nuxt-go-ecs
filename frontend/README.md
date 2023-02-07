# Nuxt 3
## packages
### ESLint
- [Nuxt 3 + Tailwind + esLint + Prettierの環境を作成する](https://zenn.dev/kigi/scraps/a2e38ff5bb0889)
- [Nuxt 3 with ESLint](https://qiita.com/kentarou_masuda/items/c0180fe383b01ba54cbf)
- [npm使用時の「: command not found」について](https://qiita.com/DaisukeNishi/items/c696860da87e559b1dbe)

@frontend dir 
$ yarn add -D eslint eslint-plugin-vue
$ yarn add -D typescript @typescript-eslint/eslint-plugin @typescript-eslint/parser @nuxtjs/eslint-config-typescript
$ npx eslint --init
> To check syntax and find problems
> JavaScript modules (import/export)
> Vue.js
> Does your project use TypeScript? Yes
> Where does your code run? Node
> What format do you want your config file to be in? JavaScript
> Would you like to install them now with npm? No

% .eslintrc.jsを修正
$ npx lint