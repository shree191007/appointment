<template>
  <v-app>


    <v-main>
      <Baseline/>
    </v-main>
  </v-app>
</template>

<script>
import Baseline from './components/Baseline';

export default {
  name: 'App',

  components: {
    Baseline
  },

  data: () => ({
    //
  }),
};
</script>
