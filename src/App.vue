<template>
  <v-app>
      <NavigationBar/>
      <v-content>
          <router-view/>
      </v-content>
      <Dialog/>
      <Bottom/>
  </v-app>
</template>

<script>
import NavigationBar from './components/NavigationBar'
import Dialog from './components/Dialog'
import Bottom from './components/Bottom'

export default {
    name: 'App',
    components:{
        NavigationBar,
        Dialog,
        Bottom,
    },
    data: () => ({
    //
    }),
};
</script>

