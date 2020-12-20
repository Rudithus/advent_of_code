load 'passport.rb'

def input
  File.open('input.txt').read.split(/\n\n/)
end

passports = input.map { |line| Passport.new(line) }

puts passports.count(&:valid?)
