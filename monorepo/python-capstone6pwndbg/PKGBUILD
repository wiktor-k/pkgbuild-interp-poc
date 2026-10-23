# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>

pkgname=python-capstone6pwndbg
# NOTE: this needs to be kept in compatibility sync with pwndbg
# https://github.com/pwndbg/pwndbg/blob/dev/pyproject.toml
pkgver=8.0.0
pkgrel=1
pkgdesc='Lightweight multi-platform, multi-architecture disassembly framework (pwndbg only)'
url='https://www.capstone-engine.org/index.html'
arch=(x86_64)
license=(BSD-3-Clause)
makedepends=(
  cmake
  git
  glibc
  python
  python-build
  python-installer
  python-setuptools
  python-wheel
)
source=(
  "git+https://github.com/pwndbg/capstone6pwndbg.git#tag=v${pkgver/Alpha/-Alpha}"
  git+https://github.com/capstone-engine/capstone.git
)
sha512sums=('326e33dbb5de3597bb3e6139a764a99a0a74d1ca065846192d26233298214d490cc2916edc8cb53458d2418f79cc4e8cbbd50b229443350fd5c80902cb11248a'
            'SKIP')
b2sums=('3429dc1c9da997b007c26943749ecec63471f1ac5b8e615abc406cf7336deb78c9df81e89bc5678b0ef2c7350b76faee8153c68ec7778339979b6155c7c030a6'
        'SKIP')

prepare() {
  cd capstone6pwndbg
  git config submodule."capstone".url "${srcdir}/capstone"
  git -c protocol.file.allow=always submodule update --recursive
  git -C capstone apply ../python-rename.patch
}

build() {
  cd capstone6pwndbg/capstone/bindings/python
  python -m build --wheel --no-isolation
}

package_python-capstone6pwndbg() {
  depends=(
    glibc
    python
  )
  cd capstone6pwndbg/capstone/bindings/python
  python -m installer --destdir="${pkgdir}" dist/*.whl
  install -Dm 644 ../../LICENSES/* -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim: ts=2 sw=2 et:
