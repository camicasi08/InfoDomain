<template>
    <div>
    <div class="row">
        <div class="col col-md-10 offset-1">
                <div class="card">
                    <div class="card-body">
                    <div class="row">
                            <div class="col col-md-2">
                                <div class="box-image">
                                    <div v-if="domain.logo != ''">
                                        <img v-bind:src="domain.logo">
                                    </div>
                                    <div v-else>
                                        <img src="../assets/internet-logo.png">
                                    </div>
                                    
                                    
                                </div>
                            </div>
                            <div class="col col-md-6 align-self-center">
                                <h4 v-if="domain.title != ''" class="text-center">{{domain.title}}</h4>
                                <h4 v-else class="text-center">{{domain.name}}</h4>
                            </div>
                            <div class="col col-md-1 align-self-center">
                                <h4><span class="badge badge-danger" title="Previous SSL Grade">{{domain.previous_ssl_grade}}</span></h4>
                            </div>
                            <div class="col col-md-1 align-self-center">
                                <h4><span class="badge badge-success" title="SSL Grade">{{domain.ssl_grade}}</span></h4>
                            </div>

                            <div class="col col-md-2">
                                <b-button v-b-toggle="domain.name" variant="link">Ver Detalles</b-button>
                            </div>
                        </div>
                    </div>
                    <b-collapse v-bind:id="domain.name" class="mt-2">
                        <div class="row">
                            <div class="col col-3">
                                 <div class="form-check">
                                     
                                    <input v-if="domain.is_down" type="checkbox" class="form-check-input" v-bind:value="domain.is_down" disabled checked>
                                    <input v-else type="checkbox" class="form-check-input" v-bind:value="domain.is_down" disabled>
                                    <label class="form-check-label">¿No Responde?</label>
                             </div>
                            </div>
                            <div class="col col-4">
                                <div class="form-check">
                                    <input v-if="domain.servers_changed" type="checkbox" class="form-check-input" v-bind:value="domain.servers_changed" disabled checked>
                                    <input v-else type="checkbox" class="form-check-input" v-bind:value="domain.servers_changed" disabled>
                                    <label class="form-check-label">¿Cambiaron los servidores?</label>
                                </div>
                            </div>
                        </div>
                       
                         
                    <b-card>
                        
                        <table class="table table-hover">
                            <thead>
                                <tr>
                                <th scope="col">Dirección</th>
                                <th scope="col">Propietario</th>
                                <th scope="col">País</th>
                                <th scope="col">Grado SSL</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr 
                                    v-for="server in domain.servers"
                                    v-bind:server="server"
                                    v-bind:key="server.address">
                                    <td> {{server.address}}</td>
                                    <td class="text-capitalize"> {{server.owner}}</td>
                                    <td class="text-uppercase"> {{server.country}}</td>
                                    <td> <span class="badge badge-success" title="SSL Grade">{{server.ssl_grade}}</span></td>
                                </tr>
                            </tbody>
                        </table>
                      
                    </b-card>
                    </b-collapse>
                
                </div>
        </div>
    </div>
    <br>
    </div>
</template>

<script>
export default {
    name: 'DoaminComponent',
    props: ['domain']
}
</script>
<style scoped>
.box-image {
    width: 50px;
    
}
img {
     max-width: 100%;
    max-height: 100%;
    display: block; /* remove extra space below image */
}
.div-centered {
  text-align: left; 
  margin: auto; 
  
}
</style>