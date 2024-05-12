<template>
  <div class="container">
    <header class="jumbotron">
      <h3>{{ content }}</h3>
    </header>
  </div>
</template>

<script>
import UserService from "../services/employee.service";

export default {
  name: "Admin",
  data() {
    return {
      content: "",
    };
  },
  mounted() {
    UserService.getAdminBoard().then(
        (response) => {
          this.content = response.data;
        },
        (error) => {
          this.loading = false;
          this.message = error.message + ": " + error.response.data.error;
        }
    );
  },
};
</script>