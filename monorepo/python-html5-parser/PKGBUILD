# Maintainer: Jelle van der Waa <jelle@archlinux.org>
# Contributor: Eli Schwartz <eschwartz@archlinux.org>

_pkgname=html5-parser
pkgname=python-html5-parser
pkgver=0.5.0
pkgrel=1
pkgdesc="Fast C based HTML 5 parsing for python"
arch=('x86_64')
url="https://github.com/kovidgoyal/${_pkgname}"
license=('Apache-2.0')
depends=('python-chardet' 'python-lxml')
makedepends=('python-setuptools' 'python-build' 'python-installer' 'python-wheel' 'git')
checkdepends=('python-beautifulsoup4')
optdepends=('python-beautifulsoup4: to use the soup treebuilder')
source=("git+${url}#tag=v${pkgver}?signed")
sha256sums=('0a3b6bbfc1855554277e4a686687d7fe609043a23a6b05caf8285029da341e3a')
b2sums=('56a6c23558e318a8a0926a8eefa7f58f0ddd4aefae376a03b5f9e42c10fa2472c6895897becb090564b215e9715ebdb710be06c445490cf00725adb28ac69aad')
validpgpkeys=('3CE1780F78DD88DF45194FD706BC317B515ACE7C') # Kovid Goyal (New longer key) <kovid@kovidgoyal.net>

build() {
    cd "${srcdir}"/${_pkgname}

    python -m build --wheel --no-isolation
}

check() {
    cd "${srcdir}"/${_pkgname}
    local python_version=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')

    # For pytest
    PYTHONPATH="$PWD/build/lib.linux-$CARCH-cpython-$python_version" python run_tests.py
}

package() {
    cd "${srcdir}"/${_pkgname}

    python -m installer --destdir="$pkgdir" dist/*.whl
}
