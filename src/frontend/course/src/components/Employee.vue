<template>
  <div class="col-md-12">
    <div class="card card-container">
      <img
          id="profile-img"
          src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
          class="profile-img-card"
       alt="Not found"/>
      <Form @submit="handleInfoCard" :validation-schema="schema">
        <div v-if="!created">
          <div class="form-group">
            <label for="phone">Номер телефона</label>
            <Field name="phone" type="text" class="form-control" />
            <ErrorMessage name="phone" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="post">Должность</label>
            <Field name="post" type="text" class="form-control" />
            <ErrorMessage name="post" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="password">Пароль</label>
            <Field name="password" type="password" class="form-control" />
            <ErrorMessage name="password" class="error-feedback" />
          </div>
          <div class="form-group">
            <button class="btn btn-primary btn-block" :disabled="loading">
              <span
                  v-show="loading"
                  class="spinner-border spinner-border-sm"
              ></span>
              Создать карточку
            </button>
          </div>
        </div>
        <div v-else>
          <div class="spinner-border spinner-border-sm">
            <label for="temp">TEMP</label>
          </div>
        </div>
      </Form>

      <div
          v-if="message"
          class="alert"
          :class="created ? 'alert-success' : 'alert-danger'"
      >
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script>
import EmployeeService from "../services/employee.service";
import * as yup from "yup";
import {Form, Field, ErrorMessage} from "vee-validate";

let ok;

export default {
  name: "Employee",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      phone: yup
          .string()
          .required("Введите номер телефона")
          .min(11, "Некорректный номер телефона"),
      post: yup
          .string()
          .required("Введите должность"),
      password: yup
          .string()
          .required("Введите пароль")
    });
    return {
      created: false,
      loading: false,
      message: "",
      schema,
    };
  },
  mounted() {
    // UserService.getHomePage().then(
    //     (response) => {
    //       this.content = response.data;
    //     },
    //     (error) => {
    //       this.content =
    //           (error.response &&
    //               error.response.data &&
    //               error.response.data.message) ||
    //           error.message ||
    //           error.toString();
    //     }
    // );
    EmployeeService.createEmployeeInfoCard().then(
        (response) => {
          this.content = response.data;
        }
    );
    EmployeeService.getEmployeeInfoCard().then(
        (response) => {
          this.content = response.data;
        }
    );
  },
  methods: {
    handleInfoCard(data) {
      if (ok === "true") {
        this.message = "OKOKOK";
      }
      else {
        this.message = "ERROR";
        ok = "true";
      }
    },
  },
};
</script>