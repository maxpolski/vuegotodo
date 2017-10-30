import Vue from 'vue';

import App from './components/App.vue';
import store from './store';

// just for eslint correctness
export default new Vue({
  el: '#app',
  store,
  render: h => h(App),
});
