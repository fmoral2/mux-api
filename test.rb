
#         list1 = [1, 2, 3, 3,7,9,4,0]
#         list2 = [1,7,6,3,9]
#         list3 = list1 + list2
#         p list3.sort.uniq!

# #######################################

#          puts "Enter a name"
#         name1 = gets.chomp.to_s
#         if name1 == name1.reverse
#             puts "The name is a palindrome"
#         else
#             puts "The name is not a palindrome"
#         end
# #######################################
        require 'httparty'
        require 'json'
        require 'rspec'




        describe 'API TEST', :t do
          it 'api test' do
     response = HTTParty.get('https://dog.ceo/api/breeds/list/all',
       :headers => { 'Accept' => 'application/json' })  

       code= expect(response.code).to eq(200)
       res = expect(response.parsed_response['message']['bulldog'][2]).to eq('french')
       
   end
end
        