<template>
    <div class="home">
        <img src="../assets/main.png" width="45%" height="45%" style="margin-top: 50px;">
        <h1>Domain Driven Design in Go</h1>
        <p>
            To get or customize this project,<br>
            check out the
            <a href="https://github.com/1ambda/domain-driven-design-go" target="_blank">repository</a>.
        </p>
        <h3>Infrastructure Layer</h3>
        <ul>
            <li><a href="https://cloud.google.com/"
                   target="_blank">GCP</a></li>
            <li><a href="https://kubernetes.io/"
                   target="_blank">Kubernetes</a></li>
            <li><a href="https://istio.io/"
                   target="_blank">Istio</a></li>
            <li><a href="https://www.spinnaker.io/"
                   target="_blank">Spinnaker</a></li>
            <li><a href="http://terraform.io/"
                   target="_blank">Terraform</a></li>
            <li><a href="https://prometheus.io/"
                   target="_blank">Prometheus</a></li>
        </ul>
        <h3>Application Layer</h3>
        <ul>
            <li><a href="https://github.com/1ambda/domain-driven-design-go/tree/master/service-gateway/internal/domain"
                   target="_blank">E-commerce Domain</a></li>
            <li><a href="https://swagger.io/tools/swagger-codegen/"
                   target="_blank">Swagger Codegen</a></li>
            <li><a href=""
                   target="_blank">Automated CI / CD Pipeline</a></li>
        </ul>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component({
    components: {
    },
})
export default class Home extends Vue {
    public $refs: any
    public $notify: any
    public $router: any
    public $store: any
}
</script>

<style scoped lang="scss">
    .home {
        text-align: center;
    }

    h3 {
        margin: 40px 0 0;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }
</style>
