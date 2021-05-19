
import Vue from "vue";
import App from './App.vue'
import VueRouter from 'vue-router'
import vuetify from './plugins/vuetify';
import routes from './router/route'
import VueResource from 'vue-resource'
import 'vue-event-calendar/dist/style.css' //^1.1.10, CSS has been extracted as one file, so you can easily update it.
import vueEventCalendar from 'vue-event-calendar'
Vue.use(vueEventCalendar, {locale: 'en'}) 
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
