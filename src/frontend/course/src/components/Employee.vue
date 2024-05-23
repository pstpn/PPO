<template>
  <div class="card">
    <div :class="['card card-container', cardValidationClass]">
      <div v-if="!created">
        <img
            id="profile-img"
            :src="imageUrl"
            sizes="(max-width:100px) 20px, 5vw"
            class="profile-img-card"
            alt="Not found"
        />
        <h5 class="card-title">Личная информация</h5>
        <Form class="card-body" ref="form" @submit="handleProfile" :validation-schema="schema">
          <div class="form-group">
            <Field type="file" name="image" class="form-control-file" accept="image/jpeg" @change="handleImageUpload" />
            <ErrorMessage name="image" class="text-danger" />
          </div>
          <div class="form-group">
            <label for="documentType">Тип документа, удостоверяющего личность</label>
            <Field name="selectedType" as="select" class="form-control">
              <option value="" disabled selected>Выберите тип документа</option>
              <option v-for="(documentType, index) in documentTypes" :key="index" :value="documentType">
                {{ documentType }}
              </option>
            </Field>
            <ErrorMessage name="selectedCompany" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="serialNumber">Серийный номер документа</label>
            <Field name="serialNumber" type="text" class="form-control" />
            <ErrorMessage name="serialNumber" class="error-feedback" />
          </div>
          <div class="form-group" v-for="(field, index) in documentFields" :key="index">
            <label :for="'documentField' + index">Поле документа {{ index + 1 }}</label>
            <div class="input-group">
              <input v-model="field.type" type="text" class="form-control" placeholder="Тип поля">
              <input v-model="field.value" type="text" class="form-control" placeholder="Значение поля">
              <div class="input-group-append">
                <button @click="removeField(index)" class="btn btn-outline-secondary" type="button">Удалить</button>
              </div>
            </div>
          </div>
          <div class="form-group">
            <button class="btn btn-primary col-md-12" :disabled="loading">
              <span v-show="loading" class="spinner-border spinner-border-sm"></span>
                Создать карточку
            </button>
            <button @click="addField" class="btn btn-dark col-md-12">Добавить поле</button>
          </div>
        </Form>
      </div>
      <div v-else>
        <img
            id="profile-img"
            :src="employeePhotoURL"
            sizes="(max-width:100px) 20px, 5vw"
            class="profile-img-card"
            alt="Not found"
        />
        <div class="card-body">
          <h5 class="card-title">Личная информация</h5>
          <p class="card-text">Тип документа: {{ profileInfo.documentType }}</p>
          <p class="card-text">Серийный номер документа: {{ profileInfo.serialNumber }}</p>
          <table class="table">
            <tbody>
            <tr v-for="(pair, index) in profileInfo.documentFields" :key="index">
              <td>{{ pair.type }}</td>
              <td>{{ pair.value }}</td>
            </tr>
            </tbody>
          </table>
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
    const schema = yup.object().shape({
      image: yup.mixed().required("Загрузите ваше фото или изображение документа, содержащего фото"),
      serialNumber: yup.string().required("Введите серийный номер документа!"),
    });
    return {
      loading: false,
      message: "",
      schema,
      documentTypes: [ "Паспорт", "Водительские права" ],
      documentFieldTypes: [ "Дата выдачи", "Выдавший орган" ],
      documentFields: [],
      imageUrl: "//ssl.gstatic.com/accounts/ui/avatar_2x.png",
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn
    },
    created() {
      if (!this.profileInfo) {
        this.$store.dispatch("employee/getProfile").then(
            () => {
              this.$store.dispatch("employee/getEmployeePhoto");
            },
            (error) => {
              if (error.response && error.response.status === 404) {
                this.$store.state.employee.profile = null;
              }
            }
        )
      }
      return this.profileInfo;
    },
    profileInfo() {
      return this.$store.state.employee.profile;
    },
    employeePhotoURL() {
      return this.$store.state.employee.photoURL;
    },
    cardValidationClass() {
      if (!this.created)
        return 'not-created';
      if (this.profileInfo == null)
        return 'card-invalid';
      return this.profileInfo.isConfirmed ? 'card-valid' : 'card-invalid';
    }
  },
  methods: {
    addField() {
      this.documentFields.push({ type: "", value: "" });
    },
    removeField(index) {
      this.documentFields.splice(index, 1);
    },
    handleImageUpload(event) {
      const file = event.target.files[0];
      if (file) {
        this.image = file;

        const reader = new FileReader();
        reader.onload = (e) => {
          this.imageUrl = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },
    handleProfile(profile) {
      if (!this.created) {
        this.$refs.form.validate().then(success => {
          if (success) {
            this.fillProfile(profile);
          }
        });
      }
    },
    fillProfile(profile) {
      let user = JSON.parse(localStorage.getItem('user'));

      const formData = new FormData();
      formData.append('image', this.image);
      formData.append('profileData', new Blob([JSON.stringify({
        serialNumber: profile.serialNumber,
        documentType: profile.selectedType,
        documentFields: this.documentFields,
      })], { type: 'application/json' }));

      this.loading = true;

      this.$store.dispatch("employee/fillProfile", formData).then(
          () => {
            this.loading = false;
            window.location.reload();
          },
          (error) => {
            this.loading = false;
            if (error.response && error.response.status === 401) {
              this.$store.dispatch('auth/refreshTokens', user).then(
                  response => {
                    this.fillProfile(profile);
                  },
                  (error) => {
                    if (error.response && error.response.status === 401) {
                      this.$store.dispatch('auth/logout');
                      this.$router.push('/login');
                    } else {
                      this.message = error.message + ": " + error.response.data.error;
                    }
                  }
              );
            } else {
              this.message = error.message + ": " + error.response.data.error;
            }
          }
      );
    },
  },
};
</script>

<style scoped>
.card-container {
  background-color: #f7f7f7;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  width: 100%;
  max-width: 1200px;
}

.profile-img-card {
  display: block;
  margin: 20px auto;
  border-radius: 50%;
  width: 250px;
  height: 250px;
  object-fit: cover;
}

.card-body {
  padding: 20px;
}

.card-title {
  font-size: 1.5em;
  margin-bottom: 10px;
  color: #333;
  text-align: center;
}

.card-text {
  font-size: 1em;
  color: #666;
  margin-bottom: 10px;
}

.card-subtitle {
  font-size: 1.1em;
  margin-top: 15px;
  margin-bottom: 10px;
  color: #333;
}

.card-valid {
  border: 2px solid green;
}

.card-invalid {
  border: 2px solid red;
}

.table {
  width: 100%;
  margin-top: 15px;
  border-collapse: collapse;
}

.table td {
  padding: 10px;
  border: 1px solid #ddd;
}

.input-group {
  width: 100%;
}

.input-group-append {
  flex: none;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
}

.form-control {
  width: 100%;
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  transition: border-color 0.3s;
}

.form-control:focus {
  border-color: #2575fc;
  outline: none;
  box-shadow: 0 0 5px rgba(37, 117, 252, 0.5);
}

.btn {
  display: inline-block;
  font-size: 1em;
  font-weight: 600;
  text-align: center;
  text-decoration: none;
  padding: 10px 15px;
  margin: 10px 5px;
  border-radius: 5px;
  transition: background-color 0.3s ease;
}

.btn-primary {
  background-color: #007bff;
  color: #fff;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.btn-dark {
  background-color: #343a40;
  color: #fff;
}

.btn-dark:hover {
  background-color: #1d2124;
}

.spinner-border {
  display: inline-block;
  width: 1.5rem;
  height: 1.5rem;
  vertical-align: text-bottom;
  border: .25em solid currentColor;
  border-right-color: transparent;
  border-radius: 50%;
  animation: spinner-border .75s linear infinite;
}

.error-feedback {
  color: #dc3545;
  font-size: 14px;
  margin-top: 4px;
}

</style>
