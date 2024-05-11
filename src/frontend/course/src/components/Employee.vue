<template>
  <div class="col-md-12">
    <div class="card card-container">
      <img
          id="profile-img"
          src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
          class="profile-img-card"
          alt="Not found"
      />
      <div v-if="!created">
        <Form @submit="handleInfoCard" :validation-schema="schema">
          <div class="form-group">
            <button class="btn btn-primary btn-block" :disabled="loading">
              <span v-show="loading" class="spinner-border spinner-border-sm"></span>
              Создать карточку
            </button>
          </div>
        </Form>
      </div>
      <div v-else>
        <div class="card-body">
          <h5 class="card-title">Информация о пользователе</h5>
          <p class="card-text">Номер телефона: {{ infoCardInfo.phone }}</p>
          <p class="card-text">Должность: {{ infoCardInfo.post }}</p>
          <!-- Здесь вы можете отобразить другие свойства пользователя -->
        </div>
      </div>
      <div v-if="message" :class="created ? 'alert-success' : 'alert-danger'" class="alert">
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "Employee",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    return {
      loading: false,
      message: "",
      schema: yup.object().shape({
      }),
    };
  },
  computed: {
    created() {
      return this.$store.state.created;
    },
    infoCardInfo() {

    },
  },
  methods: {
    handleInfoCard() {
      if (!this.created) {
        this.createEmployeeInfoCard();
      }
    },
    createEmployeeInfoCard() {
      let user = JSON.parse(localStorage.getItem('user'));

      this.loading = true;

      this.$store.dispatch("employee/createEmployeeInfoCard").then(
          () => {
            this.loading = false;
            this.$store.state.created = true;
            this.$router.push("/infocard");
          },
          (error) => {
            this.loading = false;
            if (error.response && error.response.status === 401) {
              this.$store.dispatch('auth/refreshTokens', user).then(
                  response => {
                    this.handleInfoCard();
                  },
                  (error) => {
                    if (error.response && error.response.status === 401) {
                      this.$store.dispatch('auth/logout');
                      this.$router.push('/login');
                    } else {
                      this.message = error.message || error.toString();
                    }
                  }
              );
            } else {
              this.message = error.message || error.toString();
            }
          }
      );
    },
  },
};
</script>
