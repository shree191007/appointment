<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
      clipped
      dense
      temporary
    ><v-list dense>
<v-list-item link to="/Doctor">
<v-list-item-content>
<v-list-item-title>Doctor</v-list-item-title>
</v-list-item-content>
</v-list-item>
</v-list>
<v-list dense>
<v-list-item link to="/patient">
<v-list-item-content>
<v-list-item-title>patient</v-list-item-title>
</v-list-item-content>
</v-list-item>
</v-list>
      <!--  -->
    </v-navigation-drawer>

    <v-app-bar app dense>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title dark>Appointment</v-toolbar-title>
    </v-app-bar>

<v-container fluid>
<router-view></router-view>
</v-container>

    <v-main>
      <!--  -->
    </v-main>
  </v-app>
</template>

<script>
  export default {
    data: () => ({ drawer: null }),
  }
</script>
