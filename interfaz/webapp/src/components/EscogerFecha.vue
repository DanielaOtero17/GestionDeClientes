<template>
  <v-container>
<v-row>
 <v-col>
    <v-row class = "text-center">
      <v-col cols="12"></v-col>
       <v-col class = "mb-4">
        <h1 class = "display font-weight-bold mb-2">Escoja la fecha</h1>
        <v-date-picker v-model="picker"></v-date-picker>
      </v-col>
        
    </v-row>
  </v-col>
   
   <v-col class="center">
    <div>
      <v-row class="text-center">
        <v-col class = "mb-4">
        <v-btn  color="primary" large rounded disable @click="cargarDatos"  > Cargar datos </v-btn>
        </v-col>
      </v-row>
      <v-row class="text-center">
        <v-col class = "text-center">
          <v-btn  color="secundary" large rounded  :to="{path:'/compradores'}" > Ver Compradores </v-btn>
        </v-col>
      </v-row>
    </div>
  </v-col>

</v-row>
  </v-container>
</template>

<script >
import axios from 'axios'
  export default {
    name: 'escogerfecha',
    data () {
    
      return {
        picker: new Date().toISOString().substr(0, 10),
        
      }
    },
    methods:{
    cargarDatos(){
     var tiempo = this.picker.split("-")
     var dateFormato = new Date(tiempo[0]+"/"+tiempo[1]+"/"+tiempo[2])
    var url = 'http://localhost:8090/import/' + dateFormato.getTime()
    axios.get(url);    
    }

    }
    }
</script>

<style>

.center{
  display: flex;
  align-items:center;
  justify-content:center;
}

</style>