# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>

pkgname=pwndbg
pkgver=2026.09.15
pkgrel=1
pkgdesc='Makes debugging with GDB suck less'
url='https://github.com/pwndbg/pwndbg'
arch=(any)
license=(MIT)
depends=(
  debuginfod
  gdb
  ipython
  python
  python-capstone
  python-capstone6pwndbg
  python-niche-elf
  python-psutil
  python-pt
  python-ptrace
  python-pwntools
  python-pycparser
  python-pyelftools
  python-pygments
  python-requests
  python-rich
  python-setuptools
  python-sortedcontainers
  python-tabulate
  python-typing_extensions
  python-unicorn
  # TODO:
  # decomp2dbg
  which
)
makedepends=(
  python-build
  python-hatchling
  python-installer
  python-poetry-core
  python-wheel
)
optdepends=(
  'checksec: checksec command support'
  'ropper: ropper command support'
  'ropgadget: ropgadget command support'
  'radare2: radare2 command support'
  'rizin: rizin command support'
  'one_gadget: command to find ROP one_gadget'
)
source=(
  https://github.com/pwndbg/pwndbg/archive/${pkgver}/${pkgname}-${pkgver}.tar.gz
)
sha512sums=('a699a6619936f11a9752d48d23c3ae75e048bebf9d1d59744dc6d3996c0abfd91ff181af1180f857c5ad3e321ca70d5cc92e353393001e8f9ede229a103e3a19')
b2sums=('75ddd2b6bcc81cb8f73a9b14c7d28b04f31d8358a7a5bcff1ae87f98cd5772bab9eeedea151908ab6d3f2777e894d23275df00c64c4f3c195732340caafa9ab0')

prepare() {
  cd ${pkgname}-${pkgver}
  rm -rf profiling
}

build() {
  cd ${pkgname}-${pkgver}
  python -m compileall ./*.py
  python -O -m compileall ./*.py
  python -m build --wheel --no-isolation
}

package() {
  cd ${pkgname}-${pkgver}

  python -m installer --destdir="${pkgdir}" dist/*.whl

  install -vd "${pkgdir}/usr/share/pwndbg"
  cp -r ./*.py __pycache__ "${pkgdir}/usr/share/pwndbg"

  install -vDm 644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}"
  install -vDm 644 LICENSE.md -t "${pkgdir}/usr/share/licenses/${pkgname}"
}

# vim: ts=2 sw=2 et:
