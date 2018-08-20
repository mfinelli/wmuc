lib = File.expand_path("../lib", __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require "wmuc/version"

Gem::Specification.new do |spec|
  spec.name          = "wmuc"
  spec.version       = WMUC::VERSION
  spec.authors       = ["Mario Finelli"]
  spec.email         = ["mario@finel.li"]
  spec.summary       = %q{A repository manager}
  spec.description   = %q{Clone repositories into a specific directory structure.}
  spec.homepage      = "https://github.com/mfinelli/wmuc"
  spec.license       = "GPL-3.0+"
  spec.files         = `git ls-files -z`.split("\x0").reject do |f|
    f.match(%r{^(test|spec|features)/})
  end
  spec.bindir        = 'bin'
  spec.executables = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]
  spec.add_development_dependency "bundler", "~> 1.16"
end
