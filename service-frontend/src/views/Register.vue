<template>
    <el-row type="flex" justify="center">
        <el-col :xs="16" :sm="12" :md="8" :lg="6">
            <div style="margin: 20px;"></div>
            <el-form :label-position="'right'" :rules="rules" label-width="100px" :model="ruleForm" ref="ruleForm">
                <el-form-item label="ID" prop="uid">
                    <el-input v-model="ruleForm.uid"></el-input>
                </el-form-item>

                <el-form-item label="Email" prop="email">
                    <el-input v-model="ruleForm.email"></el-input>
                </el-form-item>

                <el-form-item label="Password">
                    <el-input type="password" v-model="ruleForm.password" auto-complete="off"></el-input>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="submitForm('ruleForm')">Create</el-button>
                    <el-button @click="resetForm('ruleForm')">Reset</el-button>
                </el-form-item>
            </el-form>
        </el-col>
    </el-row>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { AuthAPI } from '@/common/auth.service.ts'
import { Exception, RegisterRequest } from '../generated/swagger'

@Component({ components: {} })
export default class Register extends Vue {
    public $refs: any
    public $notify: any
    public $router: any

    private ruleForm = {
        uid: '',
        email: '',
        password: '',
    }

    private rules = {
        uid: [
            { required: true, message: 'Please input id', trigger: 'blur' },
            { min: 4, max: 30, message: 'Length should be 4 to 30', trigger: 'blur' },
            { pattern: /^([a-zA-Z0-9]+)$/, message: 'Please use alpha numeric only', trigger: 'blur' } ],
        email: [
            { required: true, message: 'Please input email address', trigger: 'blur' },
            { type: 'email', message: 'Please input correct email address', trigger: [ 'blur', 'change' ] },
            { min: 4, max: 30, message: 'Length should be 4 to 30', trigger: 'blur' },
        ],
        password: [
            { required: true, message: 'Please input password', trigger: 'blur' },
        ],
    }

    public submitForm(formName: string) {
        this.$refs[ formName ].validate((valid: any) => {
            if (!valid) {
                this.$notify({
                    title: `Validation Failed`,
                    message: 'Please insert required values',
                    type: 'warning',
                })
                return
            }

            const request: RegisterRequest = {
                uid: this.ruleForm.uid,
                password: this.ruleForm.password,
                email: this.ruleForm.email,
            }

            AuthAPI.register(request, { credentials: 'include' })
                .then((response) => {
                    this.$notify({
                        title: 'Success',
                        message: `Created ${response.uid}`,
                        type: 'success',
                    })

                    this.$router.push('/login')
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

    public resetForm(formName: string) {
        this.$refs[ formName ].resetFields()
    }
}
</script>
