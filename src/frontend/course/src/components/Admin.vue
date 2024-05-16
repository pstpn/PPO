<template>
  <div class="search-container">
    <div class="search-bar">
      <div class="search-input-container">
        <input type="text" v-model="searchQuery" @input="searchEmployees" placeholder="Поиск сотрудников" class="search-input">
        <div class="search-input-icon">
          <i class="fas fa-search"></i>
        </div>
      </div>
      <div class="filter-dropdown" @mouseover="isDropdownVisible = true" @mouseleave="isDropdownVisible = false">
        <div class="dropdown-trigger">
          Поиск по полю
        </div>
        <transition name="fade">
          <div v-if="isDropdownVisible" class="dropdown-content">
            <select v-model="searchBy" @change="searchEmployees" class="filter-select">
              <option value="full_name">ФИО</option>
              <option value="phone_number">Номер телефона</option>
            </select>
          </div>
        </transition>
      </div>
      <div class="filter-dropdown" @mouseover="isSortDropdownVisible = true" @mouseleave="isSortDropdownVisible = false">
        <div class="dropdown-trigger">
          Направление сортировки
        </div>
        <transition name="fade">
          <div v-if="isSortDropdownVisible" class="dropdown-content">
            <select v-model="sortDirection" @change="searchEmployees" class="filter-select">
              <option value="ASC">По возрастанию</option>
              <option value="DESC">По убыванию</option>
            </select>
          </div>
        </transition>
      </div>
    </div>
    <div v-if="showSearchResults" class="search-results">
      <div v-for="employee in searchResults" :key="employee.ID" @click="viewEmployeeCard(employee.ID)" class="search-item">
        <div class="employee-info">
          <div class="employee-details">
            <div class="employee-fullName">{{ employee.FullName }}</div>
            <div class="employee-phoneNumber">{{ employee.PhoneNumber }}</div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="showEmployeeCard" class="employee-card">
      <div class="card card-container">
        <img
            id="profile-img"
            :src="employeePhotoURL"
            sizes="(max-width:100px) 20px, 5vw"
            class="profile-img-card"
            alt="Not found"
        />
        <div class="card-body">
          <h5 class="card-title">{{ selectedEmployee.FullName }}</h5>
          <p class="card-text">{{ selectedEmployee.PhoneNumber }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";

export default {
  name: "AdminSearch",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    return {
      searchQuery: '',
      searchBy: 'full_name',
      sortDirection: 'ASC',
      searchResults: [],
      showSearchResults: false,
      selectedEmployee: null,
      showEmployeeCard: false,
      isDropdownVisible: false,
      isSortDropdownVisible: false,
    };
  },
  computed: {
    employeePhotoURL() {
      return this.selectedEmployee ? this.selectedEmployee.photoURL : null;
    }
  },
  methods: {
    searchEmployees() {
      let user = JSON.parse(localStorage.getItem('user'));

      if (this.searchQuery.length > 0) {
        this.showSearchResults = true;
        const { searchQuery, searchBy, sortDirection } = this;
        console.log(searchQuery, searchBy, sortDirection);
        this.$store.dispatch('employee/getEmployees', { searchQuery, searchBy, sortDirection }).then(
            (employees) => {
              this.searchResults = employees;
            },
            (error) => {
              if (error.response && error.response.status === 404) {
                this.$store.state.employee.profile = null;
              }
              if (error.response && error.response.status === 401) {
                this.$store.dispatch('auth/refreshTokens', user).then(
                    response => {
                      this.searchEmployees();
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
        )
      } else {
        this.showSearchResults = false;
        this.searchResults = [];
      }
    },
    viewEmployeeCard(employeeId) {
      // Call backend API to fetch employee details by ID
      // Example: axios.get('/infocards/' + employeeId)
      // In this example, I'm just using the selected employee from search results
      this.selectedEmployee = this.searchResults.find(employee => employee.id === employeeId);
      this.showEmployeeCard = true; // Show employee card
    },
  }
};
</script>

<style scoped>
.search-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.search-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.search-input-container {
  position: relative;
  flex: 1;
}

.search-input {
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  width: 100%;
  transition: all 0.3s;
}

.search-input:focus {
  border-color: #007bff;
}

.search-input-icon {
  position: absolute;
  top: 50%;
  right: 10px;
  transform: translateY(-50%);
  color: #ccc;
}

.filter-dropdown {
  position: relative;
  cursor: pointer;
  width: 250px;
}

.dropdown-trigger {
  padding: 10px; /* Увеличиваем отступы для центрирования текста */
  border-radius: 5px;
  border: 1px solid #ccc;
  background-color: #007bff; /* Задаем цвет фона */
  color: #fff;
  text-align: center; /* Центрируем текст */
}

.filter-dropdown .dropdown-content {
  position: absolute;
  top: calc(100% + 5px);
  left: 0;
  background-color: #fff;
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 5px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  z-index: 10;
  opacity: 0;
  visibility: hidden;
  transition: opacity 0.3s ease, visibility 0s linear 0.3s; /* Добавляем анимацию появления и исчезания */
}

.filter-dropdown:hover .dropdown-content {
  opacity: 1;
  visibility: visible;
  transition-delay: 0s; /* Убираем задержку для появления */
}

.filter-select {
  width: 250px;
  padding: 8px;
  border-radius: 5px;
  border: none;
  background-color: #fff;
  color: #333;
  text-align: center; /* Центрируем текст */
  transition: background-color 0.3s; /* Добавляем плавный переход для цвета фона */
}

.filter-select:focus {
  outline: none; /* Убираем обводку при фокусе */
}

.search-results {
  margin-top: 20px;
}

.search-item {
  padding: 10px;
  border-bottom: 1px solid #ccc;
  cursor: pointer;
  transition: background-color 0.3s;
}

.search-item:hover {
  background-color: #f7f7f7;
}

.employee-card {
  margin-top: 20px;
}
</style>