describe('Login Page', () => {
  it('should click on button "Cadastrar" and navigate to /cadastrar', () => {
    cy.visit('http://localhost:3000/login')

    cy.get('#auth-option').click()

    cy.url().should('include', '/cadastro')
  })

  it('should click on button "Entrar"', () => {
    cy.visit('http://localhost:3000/login')

    cy.get('#login').click()

    cy.get("#email-error-label").should("exist")
    cy.get("#email-error-label").should("have.text", "o e-mail digitado deve ser válido")

    cy.get("#password-error-label").should("exist")
    cy.get("#password-error-label").should("have.text", "A senha deve contem ao menos 8 caracteres")
  })

  it('should not be able to realize login with invalid credentials', () => {
    cy.visit('http://localhost:3000/login')

    cy.get('[name="email"]').type("john.doe@test.com")
    cy.get('[name="password"]').type("12345678")

    cy.get('#login').click()

    cy.get(".login-error").should("exist")
    cy.get(".login-error").should("contain.html", "Usuário não encontrado.")
    cy.get(".login-error").should("contain.html", "Verifique os dados e tente novamente.")
  })

  it('should not be able to realize login with invalid password', () => {
    cy.visit('http://localhost:3000/login')

    cy.get('[name="email"]').type(Cypress.env("TEST_USER_EMAIL"))
    cy.get('[name="password"]').type("00000000")

    cy.get('#login').click()

    cy.get(".login-error").should("exist")
    cy.get(".login-error").should("contain.html", "E-mail ou senha inválidos.")
    cy.get(".login-error").should("contain.html", "Verifique os dados e tente novamente.")
  })

  it('should be able to change password input type from text to password', () => {
    cy.visit('http://localhost:3000/login')

    cy.get('[name="password"]').type("test")

    cy.get('#password-toggle').click()

    cy.get('[name="password"]').should("have.attr", "type", "text")
  })

  it('should be able to login', () => {
    cy.visit('http://localhost:3000/login')

    cy.log("test_user_email", Cypress.env("TEST_USER_EMAIL"))

    cy.get('[name="email"]').type(Cypress.env("TEST_USER_EMAIL"))
    cy.get('[name="password"]').type(Cypress.env("TEST_USER_PASSWORD"))

    cy.get('#login').click()

    cy.url().should('include', '/dashboard')
  })
})