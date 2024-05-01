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
            <label for="phone">Номер телефона</label>
            <Field name="phone" type="text" class="form-control" v-model="formData.phone" />
            <ErrorMessage name="phone" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="post">Должность</label>
            <Field name="post" type="text" class="form-control" v-model="formData.post" />
            <ErrorMessage name="post" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="password">Пароль</label>
            <Field name="password" type="password" class="form-control" v-model="formData.password" />
            <ErrorMessage name="password" class="error-feedback" />
          </div>
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
          <p class="card-text">Номер телефона: {{ employeeInfo.phone }}</p>
          <p class="card-text">Должность: {{ employeeInfo.post }}</p>
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
        phone: yup.string().required("Введите номер телефона").min(11, "Некорректный номер телефона"),
        post: yup.string().required("Введите должность"),
        password: yup.string().required("Введите пароль"),
      }),
      formData: {
        phone: "",
        post: "",
        password: ""
      }
    };
  },
  computed: {
    created() {
      return this.$store.state.created;
    },
    employeeInfo() {
      return this.$store.state.employeeInfo;
    },
  },
  methods: {
    handleInfoCard() {
      if (!this.created) {
        this.createEmployeeInfoCard();
      }
    },
    createEmployeeInfoCard() {
      this.loading = true;
      this.$store.dispatch("employee/createEmployeeInfoCard", this.formData).then(
          () => {
            this.loading = false;
            this.$store.state.created = true;
            this.$router.push("/infoсard");
          },
          (error) => {
            this.loading = false;
            this.message = error.message || error.toString();
          }
      );
    },
  },
};
</script>
