require 'httparty'
require 'rspec'
require 'json'

describe 'API' , :test do

  before do
    @url = 'http://localhost:8081/api/employees'
    @response = HTTParty.get(@url)
  end

  it ' should validate status code' do
  code = expect(@response.code).to eq(200)
  puts code
  end

  it ' should validate value and data type' do
    expect(@response.parsed_response[1]['name']).to eq('chico')
     expect(@response.parsed_response[1]['wage']).is_a? Float
  end

  it'should add a new emp' do
    
    body = {
      name: 'luiz',
      wage: 100.00,
      role: 'qa',
    }.to_json

    @addUser = HTTParty.post(@url , body: body,
       headers: { 'Content-Type' => 'application/json' })
    expect(@addUser.code).to eq(201)
    @getUserCreated = HTTParty.get(@url + '/' + @addUser.parsed_response['id'])
    expect(@getUserCreated.parsed_response['id']).to eq('luiz')
  end

end


describe('web test ', () => {

  it('go to page ', () => {

    cy.visit("https://www.google.com/");
    // go to search page and press enter

    cy.get('input[name="q"]').type("cypress{enter}");
    // get h3 
    
    cy.get('h3').should('contain', 'cypress').eq(10);

  });

});
