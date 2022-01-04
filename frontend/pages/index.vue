<template>
  <div>
    <Navbar />
    <div>
      <b-jumbotron header="Project Atlas" lead="Basic workout tracking">
        <p>Create an accout to start using</p>
        <b-button variant="primary" href="#">More Info</b-button>
      </b-jumbotron>
    </div>

    <div>
      <UserCreate />
    </div>

    <Footer />
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "IndexPage",
});
</script>

<style>
.div {
  margin: 0.1em;
}
</style>
