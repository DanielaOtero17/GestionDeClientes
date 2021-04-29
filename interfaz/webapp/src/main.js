import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import VueRouter from 'vue-router'

// importo los componentes
import escogerfecha from './components/EscogerFecha';
import listaCompradores from './components/ListarCompradores';


//ahora, creo los componentes
Vue.component('escogerfecha',escogerfecha);
Vue.component('listaCompradores',listaCompradores);



//uso de vue-router
Vue.use(VueRouter);

//definimos las rutas
const routes = [
  {path:'/',component:escogerfecha},
  {path: '/escogerFecha',component:escogerfecha},
  {path: '/compradores', component:listaCompradores},

]

//creamos el objeto router
const router = new VueRouter({
  routes,
  mode: 'history'
})

Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
