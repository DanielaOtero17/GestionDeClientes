<template>
<v-app>
  <v-card
    class="mx-auto overflow-hidden"
    height=100%
    width=100%
  >
    <v-system-bar color="deep-purple darken-3"></v-system-bar>

    <v-app-bar
      color="deep-purple accent-4"
      dark
      prominent
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>COMPRAS POR MONTÓN</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-filter</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>
    </v-app-bar>

    <v-navigation-drawer
      v-model="drawer"
      absolute
      bottom
      temporary
    >
      <v-list
        nav
        dense
      >
        <v-list-item-group
          v-model="group"
          active-class="deep-purple--text text--accent-4"
        >
            <v-list-item :to="{path:'/escogerFecha'}">
            <v-list-item-title>Principal</v-list-item-title>
          </v-list-item>

          <v-list-item :to="{path:'/compradores'}">
            <v-list-item-title>Compradores</v-list-item-title>
          </v-list-item>

        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>
 
     <router-view></router-view>
  </v-card>
</v-app>
</template>

<script>

  export default {
    components:{
     
    },
    data: () => ({
      drawer: false,
      group: null,
    }),

    watch: {
      group () {
        this.drawer = false
      },
    },
  }
</script>
