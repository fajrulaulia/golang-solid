# Liskov Substitution Principle (LSP)

Prinsip ini berlaku untuk hierarki pewarisan dan hanya merupakan perpanjangan dari prinsip OCP.
Artinya kita harus memastikan bahwa kelas turunan baru memperluas kelas dasar tanpa mengubah perilaku aslinya. Pada dasarnya, kelas turunan tidak boleh kurang dari kelas dasarnya.
Jika subtipe dari supertipe melakukan sesuatu yang tidak diharapkan oleh klien supertipe, maka ini merupakan pelanggaran LSP. Bayangkan sebuah kelas turunan memberikan pengecualian yang tidak diberikan oleh superkelas, atau jika kelas turunan memiliki beberapa efek samping yang tidak terduga. Kita harus mempertimbangkan bagaimana program klien menggunakan hierarki kelas. Terkadang pemfaktoran ulang kode diperlukan untuk memperbaiki pelanggaran LSP yang teridentifikasi
