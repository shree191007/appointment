import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import vuetify from './plugins/vuetify';
import routes from './router/route'
import VueResource from 'vue-resource'
Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(VueResource)
Vue.http.options.crossOrigin=true
Vue.http.headers.common['Access-Control-Allow-Origin'] = '*'
const router=new VueRouter({routes})
new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
