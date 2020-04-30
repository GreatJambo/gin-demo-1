<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                id="input-1"
                v-model="$v.user.name.$model"
                type="text"
                required
                placeholder="输入用户名称"
                :state="validateState('name')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('name')">
                用户名长度必须是2-12位
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group label="电话">
              <b-form-input
                id="input-1"
                v-model="$v.user.telephone.$model"
                type="number"
                required
                placeholder="输入手机号码"
                :state="validateState('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                手机号不符合要求
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group label="密码">
              <b-form-input
                id="input-1"
                v-model="$v.user.password.$model"
                type="password"
                required
                placeholder="输入用户密码"
                :state="validateState('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码长度必须在6-18之间
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group>
              <b-button
                @click="register"
                variant="outline-primary"
                block=""
              >注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength, maxLength } from 'vuelidate/lib/validators';

import customValidator from '@/helper/validator';
import storageService from '@/service/storageService';
import userService from '@/service/userService';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      name: {
        required,
        minLength: minLength(2),
        maxLength: maxLength(12),
      },
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
        maxLength: maxLength(18),
      },
    },
  },
  methods: {
    validateState(name) {
      // 这里是es6的 析构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      console.log('正在请求API');
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        console.log('API拒绝请求');
        return;
      }
      console.log('正在发送请求');
      // 请求api
      userService.register(this.user).then((res) => {
        // 保存token
        storageService.set(storageService.USER_TOKEN, res.data.data.token);
        userService.info((response) => {
          // 保存用户信息
          storageService.set(storageService.USER_INFO, JSON.stringify(response.data.data.user));
          // 跳转主页
          this.$router.replace({ name: 'Home' });
        });
      }).catch((err) => {
        console.log(err);
        console.log('回调失败');
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });

      console.log('register');
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
