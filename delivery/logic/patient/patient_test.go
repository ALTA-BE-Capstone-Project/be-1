package patient

// func TestValidationRegexPatient(t *testing.T) {
// 	t.Run("error user name", func(t *testing.T) {
// 		var req Req

// 		req.UserName = "hotaru123   "

// 		var l = New()

// 		var err = l.ValidationRegexPatient(req)

// 		assert.NotNil(t, err)
// 		log.Info(err)
// 	})
// 	t.Run("error name", func(t *testing.T) {
// 		var req Req

// 		req.Name = "hotaru123   "

// 		var l = New()

// 		var err = l.ValidationRegexPatient(req)

// 		assert.NotNil(t, err)
// 		log.Info(err)
// 	})

// 	t.Run("error address", func(t *testing.T) {
// 		var req Req

// 		req.Address = "hotaru123   *(&*&**&*&*&(*&(*&&^"

// 		var l = New()

// 		var err = l.ValidationRegexPatient(req)

// 		assert.NotNil(t, err)
// 		log.Info(err)
// 	})

// 	t.Run("error nik", func(t *testing.T) {
// 		var req Req

// 		req.Nik = "123456789123456a"

// 		var l = New()

// 		var err = l.ValidationRegexPatient(req)

// 		assert.NotNil(t, err)
// 		log.Info(err)
// 	})

// 	t.Run("error date", func(t *testing.T) {
// 		var req Req

// 		req.Dob = "123456789123456a"

// 		var _, err = req.ToPatient()

// 		assert.NotNil(t, err)
// 		log.Info(err)
// 	})

// 	t.Run("succeess date", func(t *testing.T) {
// 		var req Req

// 		req.Dob = "05-05-2000"

// 		var _, err = req.ToPatient()

// 		assert.Nil(t, err)
// 		log.Info(err)
// 	})

// 	t.Run("success", func(t *testing.T) {
// 		var req Req

// 		var l = New()

// 		var err = l.ValidationRegexPatient(req)

// 		assert.Nil(t, err)
// 		log.Info(err)
// 	})
// }