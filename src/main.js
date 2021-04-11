import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import vuetify from './plugins/vuetify';
import routes from './router/route'

Vue.config.productionTip = false
Vue.use(VueRouter)
const router=new VueRouter({routes})
new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
