var ckeditor

ClassicEditor
        .create( document.querySelector( '#editor' ), {
          
          toolbar: ['fontfamily', 'fontsize', 'fontcolor', 'bold', 'italic', 
          'bulletedList', 'indent', 'outdent', 'numberedList', 'link', 'blockQuote', 
          'insertTable','code', 'codeblock', 'imageinsert', 'mediaembed', 'undo', 'redo' ],
          fontFamily: {
            options: [
              'ឧត្តមាន​ជ័យ, OdorMeanChey', 'អក្សរដៃ, HandWriting',
              'គូលេន, Koulen', 'ក្រូច​ឆ្នារ, Limonf3',
              'បាយ័ន, Bayon', 'ក្រសាំង, Rooster',
              'មូល, Moul',
    				  'Arial, Helvetica, sans-serif',
 				      'Courier New, Courier, monospace',
 				      'Georgia, serif',
 				      'Lucida Sans Unicode, Lucida Grande, sans-serif',
 				      'Tahoma, Geneva, sans-serif',
 				      'Times New Roman, Times, serif',
 				      'Trebuchet MS, Helvetica, sans-serif',
				      'Verdana, Geneva, sans-serif',
            ],
            supportAllValues: true
          },
          
          fontSize: {
            options: [
                9,
                11,
                13,
                'default',
                17,
                19,
                21
            ],
            supportAllValues: true
          },
        })
        .then( editor => {
          ckeditor = editor
          editor.editing.view.focus()
          editor.editing.view.change(writer=>{
              writer.setStyle('height', '270px', editor.editing.view.document.getRoot())
          })
        })
        .catch( err => {
          console.error( err.stack )
        });