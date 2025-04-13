describe('Navigation', () => {
  it('should click on button "Conhecer Agora"', () => {
    cy.visit('http://localhost:3000/')

    cy.get('#home-btn').click()

    cy.url().should('include', '/login')
  })
})