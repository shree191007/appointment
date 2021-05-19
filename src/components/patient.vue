<template>
<v-app>
  <v-data-table
    :headers="Headers"
    :items="patients"
    align="center">
  <template v-slot:[`item.actions`]="{ item }">
            <v-icon small class="mr-2" @click="editpatient(item.id)">mdi-pencil</v-icon>
            <v-icon small @click="deletepatient(item.id)">mdi-delete</v-icon>
            <v-icon small link-to="/patients">mdi-calendar</v-icon>
          </template>
  </v-data-table>
  <v-row justify="center">
    <v-btn
      color="primary"
      dark
      @click.stop="adddialog = true"
    >
      add patient
    </v-btn>
    <v-dialog
      v-model="adddialog"
      id="adddialog"
      name="adddialog"
    >
      <v-card>
        <v-card-title class="headline">
          Add patient
        </v-card-title>
        <v-card-text>
          fill the patient_form to add a patient
        </v-card-text>
      <form class="form" v-on:submit.prevent="createpatient">
        <v-text-field

                  id="name"
                  label="name"
                  v-model="patient_form.name"
                  required
                >
        </v-text-field>
        <v-text-field
                  id="age"
                  label="age"
                  v-model="patient_form.age"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="phone"                  
                  label="phone"
                  v-model="patient_form.phone"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="diagnosis"
                  label="diagnosis"
                  v-model="patient_form.diagnosis"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="mail_id"
                  label="mail_id"
                  v-model="patient_form.mail_id"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="phone"
                  label="phone"
                  v-model="patient_form.phone"
                  required
                >
        </v-text-field>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="green darken-1"
            text
            @click="add_edit_patient()"
          >
            submit
          </v-btn>
          <v-btn
          @click="adddialog = false;
          id ='add'
            ">
            close
          </v-btn>
        </v-card-actions>
      </form>
      </v-card>
    </v-dialog>
   
  </v-row>
</v-app>
</template>
<script>
import patientservices from "@/services/patientService" 
import axios from 'axios'
axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
Host="127.0.0.1"
Port="5000"


  export default {
    data :() => ({

        patient_form :{
          name:'',
          age:'',
          diagnosis:'',
          mail_id:'',
          phone:'',
        }
      ,
      adddialog: false,
        Headers:[
            {text:'patient id', value:"id"},
            {text:'patient name', value:"name"},
            {text:'phone', value:"phone"},
            {text:'age', value:"age"},
            {text:'diagnosis', value:"diagnosis"},
            {text:'email_id', value:"mail_id"},
            {text:'actions',value:"actions"},
        ],
        patients:[] ,
    }),
    created(){
        console.log('inside created')
        this.initialize()
    },
    methods:{
        initialize()
        {this.loadpatients()},
        createpatient(){
          this.patient_form={
          name:'',
          age:'',
          diagnosis:'',
          mail_id:'',
          phone:'',
          hospital_id:'55555',
        }
          this.adddialog= true
        },
        loadpatients(){
            this.patients=[]
            patientservices.getAllpatients()
            patientservices.getAllpatients().then(item=>{
                    item.patient.forEach(
                        patient=>{this.patients.push(patient)}
                    )
            });
        },
        add_edit_patient(){
          if (this.patient_form.id==null){
          axios.post('http://localhost:5000/patients', this.patient_form)
          .then((res)=>{
            console.log(res);
            if (res.status==200){
              this.adddialog = false
              this.loadpatients()
            }
          })
          .catch((err)=>{
            console.log(err);
          })
          .finally(()=>{})}
          else{
            axios.put(`http://localhost:5000/patients/${this.patient_form.id}`,this.patient_form)
            this.loadpatients()
            this.adddialog =false
          }
        },
        editpatient(id){
          axios.get(`http://localhost:5000/patients/${id}`)
          .then((res)=>{
             console.log(res.data.patient);
             this.patient_form=res.data.patient
             this.adddialog = true
          }
)
          .catch(err => console.log(err))
        },
        deletepatient(id){
          axios.delete(`http://localhost:5000/patients/${id}`)
          this.loadpatients()
        }
    }}
</script>

  