<template>
    <v-container>
        <v-row class="text-center">
            <v-col class="mb-2">
                
            </v-col>
        </v-row>
  <v-row >
    <v-col  >
      <h2 class= "text-center elevation-1 font-wight-bold mb-3">LISTA DE CLIENTES</h2>
        <v-simple-table 
        fixed-header
        class="elevation-1"
        height="900px">
            <template v-slot:default>
                <thead>
                    <tr>
                        <th class="text-left">Identificacion</th>
                        <th class="text-left">Nombre</th>
                        <th class="text-left">Edad</th>
                        <th class="text-left">Ver info</th>
                    </tr>
                </thead>
                <tbody>
                     <tr v-for="comprador in compradores" :key="comprador.id">
                    <td class="text-left">{{comprador.id}}</td>
                    <td>{{comprador.name}}</td>
                    <td>{{comprador.age}}</td>
                    <td > 
                        <v-btn fab small color="primary" @click="cargarTransacciones(comprador.id)"><v-icon>mdi-eye</v-icon>
                        </v-btn>
                    </td>
                </tr>
                </tbody>            
                </template>
        </v-simple-table>
    </v-col>
    <v-col>
        <v-div>
<v-row class = "ajustar">  
    <v-col >
      <h2 class= "text-center elevation-1 font-wight-bold mb-3">Transacciones</h2>
    <v-simple-table fixed-header height="400px" class="elevation-1">
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">Id Transaccion</th>
          <th class="text-center">Ver productos comprados</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="transaccion in transacciones" :key="transaccion.id_trans" >
          <td>{{ transaccion.id_trans }}</td>
          <td class="text-center"> 
            <v-btn  fab small color=orange @click="cargarProductos(transaccion.id_trans)"><v-icon>mdi-magnify</v-icon>
            </v-btn>
          </td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</v-col>

<v-col >
  <h2 class= "text-center elevation-1 font-wight-bold mb-3">Productos comprados</h2>
    <v-simple-table fixed-header height="400px" class="elevation-1">
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">Nombre</th>
          <th class="text-left">Precio</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="producto in productos"
          :key="producto.name_p"
        >
          <td>{{ producto.name_p }}</td>
          <td>{{ producto.price }}</td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</v-col>

</v-row>
<v-row>
<v-col >
  <h2 class= "text-center elevation-1 font-wight-bold mb-3">Otros compradores con la misma ip</h2>
    <v-simple-table fixed-header height="400px" class="elevation-1">
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">Id</th>
          <th class="text-left">Nombre</th>
          <th class="text-left">Edad</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="otro in otros"
          :key="otro.id"
        >
      
          <td>{{ otro.id }}</td>
          <td>{{ otro.name }}</td>
          <td>{{ otro.age }}</td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</v-col>
<v-col >
  <h2 class= "text-center elevation-1 font-wight-bold mb-3">Productos recomendados</h2>
    <v-simple-table fixed-header height="400px" class="elevation-1">
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">Nombre</th>
          <th class="text-left">Precio</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="recomendado in recomendados"
          :key="recomendado.name_p"
        >
          <td>{{ recomendado.name_p }}</td>
          <td>{{ recomendado.price }}</td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</v-col>
</v-row>

</v-div>
    </v-col>

  


</v-row>

    </v-container>
</template>

<script>
import axios from 'axios'
export default {
    name:'listaCompradores',

 data(){
    
     return{
         compradores:[],
         transacciones:[],
         productos:[],
         otros:[],
         recomendados:[],
         posiciones: [],
         idResponse: ""
     }
 }
 ,
 mounted(){
    this.cargarCompradores()
  },
methods:{
    cargarCompradores(){
    axios.get('http://localhost:8090/buyers').then((response)=>{
        
        for(let i=0; i<response.data['listadoCompradores'].length; i++){
          let converted={id:response.data['listadoCompradores'][i].id, name:response.data['listadoCompradores'][i].name,age: response.data['listadoCompradores'][i].age}
          this.compradores.push(converted)
          localStorage.setItem(converted.id, JSON.stringify(converted));
        }
    }).catch(e=>{
        console.log(e)
    })
    },
    cargarTransacciones(id){
        console.log("el id es "+id)
        
        var ruta = 'http://localhost:8090/buyers/'+id
        this.idResponse = id
  
         if(this.transacciones.length>0){
              this.transacciones=[]
          }
           if(this.productos.length>0){
        this.productos=[]
      }
       if(this.otros.length>0){
        this.otros=[]
      }
      if(this.recomendados.length>0){
        this.recomendados=[]
      }
      if(this.posiciones.length>0){
        this.posiciones=[]
      }
        axios.get(ruta).then((response)=>{
          
       console.log(response.data["InformacionComprador"][0]['historialCompras'])

          for(let i=0; i<response.data["InformacionComprador"][0]['historialCompras'].length; i++){
            let converted={id_trans: response.data["InformacionComprador"][0]['historialCompras'][i]["Transaccion"].id_trans}   
            this.posiciones.push(i+"-"+converted.id_trans)         
            this.transacciones.push(converted)
          }
        
        }).catch(e=>{
          console.log(e)
        })

        this.cargarOtrosCompradores(ruta)
        this.cargarRecomendados(ruta)
    },
    cargarProductos(id_trans){

      var ruta = 'http://localhost:8090/buyers/'+this.idResponse

      for(let i=0; i<this.posiciones.length;i++){
        var itemsito = this.posiciones[i].split("-")

        if(itemsito[1]==id_trans){
          
          var posicionEscogida = itemsito[0]
          break
        }        
      }

      if(this.productos.length>0){
        this.productos=[]
      }
      axios.get(ruta).then((response)=>{

        for(let i=0; i<response.data["InformacionComprador"][0]['historialCompras'][posicionEscogida]["ProductosComprados"].length; i++){
            let converted={name_p: response.data["InformacionComprador"][0]['historialCompras'][posicionEscogida]["ProductosComprados"][i].name_p,
                          price: response.data["InformacionComprador"][0]['historialCompras'][posicionEscogida]["ProductosComprados"][i].price}   
                
            this.productos.push(converted)
        }
       }).catch(e=>{
          console.log(e)
        })

    },
    cargarOtrosCompradores(ruta){

      axios.get(ruta).then((response)=>{

      for(let i=0; i<response.data["InformacionComprador"][1]['otrosCompradores'].length; i++){
            let converted={id: response.data["InformacionComprador"][1]['otrosCompradores'][i].id,
                          name: response.data["InformacionComprador"][1]['otrosCompradores'][i].name,
                          age: response.data["InformacionComprador"][1]['otrosCompradores'][i].age}   
                
            this.otros.push(converted)
        }     

    }).catch(e=>{
      console.log(e)
    })
},
  cargarRecomendados(ruta){
  
  axios.get(ruta).then((response)=>{

      for(let i=0; i<response.data["InformacionComprador"][2]['ProductosRecomendados'].length; i++){
            let converted={name_p: response.data["InformacionComprador"][2]['ProductosRecomendados'][i].name_p,
                          price: response.data["InformacionComprador"][2]['ProductosRecomendados'][i].price} 
                
            this.recomendados.push(converted)
        }     

    }).catch(e=>{
      console.log(e)
    })

  }
}
}
</script>

<style>

.elevation-3, .ajustar{
  display: flex;
}

</style>