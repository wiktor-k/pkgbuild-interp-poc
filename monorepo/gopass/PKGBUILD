# Maintainer: Morten Linderud <foxboron@archlinux.org>
# Maintainer: Christian Rebischke <chris.rebischke@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=gopass
pkgver=1.17.4
pkgrel=1
pkgdesc="The slightly more awesome standard unix password manager for teams."
arch=('x86_64')
url="https://github.com/gopasspw/gopass"
license=('MIT')
depends=('glibc')
makedepends=('go' 'git')
optdepends=('xdotool: for typing passwords selected by dmenu'
            'xsel: clipboard support'
            'xclip: clipboard support'
            'wl-clipboard: clipboard support (wayland)')
options=("!lto")
source=("git+${url}.git#tag=v${pkgver}?signed")
sha512sums=('c9d028f7c986a3850cb3b0fb43efa6cde3eb51f9eb94692a0f63b170e2c0017ab90af8fd2dfaf0b751c16eddd48a89790fa3c5425460d795b9363932f1869325')
validpgpkeys=('7BCE61F400F3CD911EF86EF1D31FA054C85AEFAC') # Dominik Schulz <dominik.schulz@gmail.com>

prepare() {
  cd ${pkgname}
  GOFLAGS="-mod=readonly" go mod vendor -v
}

build() {
  cd ${pkgname}
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -mod=vendor -modcacherw -buildvcs=false"
  export GOPATH="${srcdir}"

  local ld_flags=" \
    -X main.version=${pkgver} \
    -X main.commit=$(git rev-parse --short HEAD) \
    -X main.date=NOTSET \
    -compressdwarf=false \
    -linkmode=external \
  "
  go build -v -ldflags "${ld_flags}"

  ./gopass completion bash > bash.completion
  ./gopass completion fish > fish.completion
  ./gopass completion zsh > zsh.completion
  go run helpers/man/main.go > gopass.1
}

check() {
  cd ${pkgname}
  export GIT_CONFIG_GLOBAL="${PWD}/gitconfig"
  git config --global user.email "builduser@archlinux.org"
  git config --global user.name "Build User"

  # Suppress GPG "insecure memory" warning that pollutes test output.
  # Exclude gpgconf tests as they test the GOPASS_GPG_OPTS parsing itself.
  # Exclude config tests as TestEnvVarsInDocs is a docs linting test (upstream issue).
  local unit_tests=$(
    go list ./... | grep -v \
      -e 'github.com/gopasspw/gopass/internal/backend/crypto/gpg/gpgconf' \
      -e 'github.com/gopasspw/gopass/internal/config'
  )
  # shellcheck disable=SC2086
  GOPASS_GPG_OPTS="--no-secmem-warning" GOPASS_BINARY=$PWD/gopass \
    go test -v ${unit_tests}
}

package() {
  cd ${pkgname}
  install -vDm755 -t "${pkgdir}/usr/bin" gopass
  install -vDm644 -t "${pkgdir}/usr/share/licenses/${pkgname}" LICENSE
  install -vDm644 -t "${pkgdir}/usr/share/man/man1" gopass.1

  install -vDm644 bash.completion "${pkgdir}/usr/share/bash-completion/completions/gopass"
  install -vDm644 fish.completion "${pkgdir}/usr/share/fish/vendor_completions.d/gopass.fish"
  install -vDm644 zsh.completion "${pkgdir}/usr/share/zsh/site-functions/_gopass"
}
