require 'optparse'

require 'wmuc/version'

module WMUC
  class CLI
    def self.parse!(args)
      OptionParser.new do |parser|
        parser.on("-h", "--help", "Prints this help") do
          puts parser
          exit
        end

        parser.on("--version", "Show the version") do
          puts "WMUC v#{WMUC::VERSION}"
          exit
        end
      end.parse!
    end
  end
end
