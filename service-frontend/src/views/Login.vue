<template>
    <el-row type="flex" justify="center">
        <el-col :xs="16" :sm="12" :md="8" :lg="6">
            <div style="margin: 40px;"></div>
            <el-form :label-position="'right'" :rules="rules" label-width="100px" :model="ruleForm" ref="ruleForm">
                <el-form-item label="ID" prop="uid">
                    <el-input v-model="ruleForm.uid"></el-input>
                </el-form-item>
                <el-form-item label="Password">
                    <el-input type="password" v-model="ruleForm.password" auto-complete="off"></el-input>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="submitForm('ruleForm')">Login</el-button>
                </el-form-item>
            </el-form>
        </el-col>
    </el-row>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapActions, mapGetters, mapMutations, mapState } from 'vuex'
import { Exception, LoginRequest } from '@/generated/swagger'
import { AuthAPI } from '@/common/auth.service.ts'

@Component({
    components: {},
    computed: {
        ...mapState([ 'uid' ]),
        ...mapGetters([ 'authenticated' ]),
    },
})
export default class Login extends Vue {
    public $refs: any
    public $notify: any
    public $router: any
    public $store: any

    public ruleForm = {
        uid: '',
        password: '',
    }

    public rules = {
        uid: [
            { required: true, message: 'Please input id', trigger: 'blur' },
            { min: 4, max: 30, message: 'Length should be 4 to 30', trigger: 'blur' },
            { pattern: /^([a-zA-Z0-9]+)$/, message: 'Please use alpha numeric only', trigger: 'blur' } ],
        password: [
            { required: true, message: 'Please input password', trigger: 'blur' } ],
    }

    public mounted() {
        if (this.$store.state.uid !== '') {
            this.$router.push('/')
            return
        }
    }

    public submitForm(formName: string) {
        this.$refs[ formName ].validate((valid: any) => {
            if (!valid) {
                this.$notify.warn({
                    title: `Validation Failed`,
                    message: 'Please insert required values',
                })
                return
            }

            const request: LoginRequest = {
                uid: this.ruleForm.uid,
                password: this.ruleForm.password,
            }

            AuthAPI.login(request, { credentials: 'include' })
                .then((response) => {
                    console.log(response)
                    if (!response.uid) {
                        return
                    }

                    this.$store.commit('login', response.uid)
                    this.$router.push('/')
                })
                .catch((response) => {
                    if (!response.json) {
                        this.$notify.error({
                            title: `Error (Connection)`,
                            message: "Server is not available",
                        })

                        return
                    }

                    response.json().then((parsed: Exception) => {
                        this.$notify.error({
                            title: `Error (${parsed.type})`,
                            message: parsed.message,
                        })
                    })
                })
        })
    }
}
</script>
