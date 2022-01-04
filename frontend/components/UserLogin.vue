<template>
  <div>
    <b-form id="createForm" @submit="onSubmit">
      <b-form-group
        id="input-group-1"
        label="Username"
        label-for="input-1"
        description="Provide your username"
      >
        <b-form-input
          id="input-1"
          v-model="form.username"
          type="text"
          placeholder="Username"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group
        id="input-group-3"
        label="Password"
        label-for="input-3"
        description="Provide your password"
      >
        <b-form-input
          id="input-3"
          v-model="form.password"
          type="password"
          placeholder="Password"
          required
        ></b-form-input>
      </b-form-group>
      <b-button type="submit" variant="primary">Login</b-button>
    </b-form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
    };
  },
  methods: {
    onSubmit(event) {
      event.preventDefault();
      try {
        await this.$auth.loginWith("local", {
          form,
        });

        this.$router.push("/");
      } catch (e) {
        this.error = e.response.data.message;
        console.log(this.error);
      }
    },
  },
};
</script>

<style>
#createForm {
  margin: 1em;
}
</style>
