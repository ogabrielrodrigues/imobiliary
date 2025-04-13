describe('Navigation', () => {
  it('should click on button "Conhecer Agora"', () => {
    cy.visit('http://localhost:3000/')

    cy.get('#know-now').click()

    cy.url().should('include', '/login')
  })

  it('should click on button "Login"', () => {
    cy.visit('http://localhost:3000/')

    cy.get('#login').click()

    cy.url().should('include', '/login')
  })

  it('should click on button "Cadastro"', () => {
    cy.visit('http://localhost:3000/')

    cy.get('#sign').click()

    cy.url().should('include', '/cadastro')
  })
})