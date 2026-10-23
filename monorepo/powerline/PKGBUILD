# Maintainer: Jiachen Yang <farseerfc@gmail.com>
# Maintainer: Robin Candau <antiz@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# AUR Maintainer: Swift Geek
# AUR Maintainer: Timothée Ravier <tim@siosm.fr>
# AUR Maintainer: Stefan Tatschner <rumpelsepp@sevenbyte.org>
# Contributor: Swift Geek
# Contributor: Pablo Olmos de Aguilera C. pablo+aur at odac dot co

pkgbase=powerline
pkgname=(
  powerline
  powerline-fonts
  vim-powerline
)
pkgver=2.8.4
pkgrel=4
pkgdesc="Statuslines and prompts for several applications including vim, zsh, bash, tmux, IPython, Awesome, i3 and Qtile"
arch=(x86_64)
url="https://github.com/powerline/powerline"
license=(MIT)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-sphinx
  python-wheel
)
checkdepends=(
  python-netifaces
  python-pexpect
  python-psutil
)
source=(
  "git+$url.git#tag=$pkgver"
  # This patch set is maintained here:
  # https://github.com/carlsmedstad/powerline/tree/archlinux
  "0001-Compile-powerline-executable-without-setuptools.patch::$url/commit/ab071647568255b7c80c776beaaf467168b4f5ec.patch"
  "0002-Point-shebangs-to-system-Python.patch::$url/commit/6324b25a2a16d5c0bf5dd05ba1bc5448215ad969.patch"
  "0003-Remove-failing-tests.patch::$url/commit/b8772a7583e25f798e4790e550e101f9ce78a2f3.patch"
  "0004-Fix-argparse-nested-argument-groups-for-Python-3.14.patch::$url/commit/9528eec9c0d79d9ee0fe8029519aa1e39220ab60.patch"
  "0005-Replace-codecs.open-with-open.patch::$url/commit/ec357c25676b0fda099fba975109f18f2930ba2d.patch"
)
b2sums=('8a288cc64c2568868550a6de18ecede51f811582456450ae169f806256ad09555c068302fe2982711e47fe1899b64ae3f3084bd5c8ca1b364a36630709af6bed'
        '1c0bd34fbfeaa87cd5decc459bb58da2ad71cb438b2d2572217f2f31c70cf59150676c896a0b0bb9f6e33f21f4c91744b5230a788f86c47636ec571e73fa1f7b'
        'c5a80e823abecfddf1ca7b4299b4fc05b2af6d3d90e87d0508566a43405757b1c99ab2cae5a68f5ae353f06310c2c91a0bfbb04b99eb5ed644bb3f103e8c7020'
        'a81c18629808868d37ef12f70558e685e0c13d1a1c157d5dc52dc81e364a61df42df1e80000c6ad2d07e049c2f3d2973375894e26e857cf1b5b15063ea1a048f'
        '4c0775a2131a22ba460538335da5cb0b3c42cc7bd7e6041edee3d28e4158c4fac33f7684155fd26d0a902fcc25f47334cccd9c19b5ba5fc57a8c00c6b6f701fe'
        '46361b62bbe50733091b52244fe23fde37d540c07c090a600f6d7088ef4dc8a8d35d1572b8e2a6936de28196f608bf5664d91ae75b533d29ed810095b893f099')

prepare() {
  cd $pkgbase
  for patch in ../*.patch; do
    patch -Np1 < "$patch"
  done
}

build() {
  cd $pkgbase
  # shellcheck disable=SC2086
  gcc $CFLAGS $LDFLAGS -o scripts/powerline client/powerline.c
  python -m build --wheel --no-isolation
  make -C docs man
}

check() {
  cd $pkgbase
  python -m unittest discover -vs tests/test_python
}

package_powerline() {
  depends=(
    glibc
    python
  )
  optdepends=(
    'powerline-fonts: patched fonts for powerline'
    'vim-powerline: vim plugin'
    'python-netifaces: for displaying internal IP addresses'
    'python-psutil: improved system information'
    'python-pygit2: improved git support'
    'which: fish bindings'
  )
  provides=(
    python-powerline
    powerline-common
  )
  replaces=(
    powerline-common
    python-powerline
  )

  cd $pkgbase
  python -m installer --destdir="$pkgdir" dist/*.whl

  install -vdm755 "$pkgdir/usr/share/powerline"
  ln -vs "$(python -c 'import site; print(site.getsitepackages()[0])')/powerline/bindings" \
    "$pkgdir/usr/share/powerline/bindings"

  install -vDm644 -t "$pkgdir/usr/lib/systemd/user" \
    powerline/dist/systemd/powerline-daemon.service
  install -vDm644 -t "$pkgdir/usr/share/man/man1" docs/_build/man/*
  install -vDm644 -t "$pkgdir/usr/share/licenses/powerline" LICENSE
}

package_powerline-fonts() {
  pkgdesc="Patched fonts for powerline"
  provides=(otf-powerline-symbols)

  cd $pkgbase
  install -vDm644 -t "$pkgdir/usr/share/fonts/OTF" font/PowerlineSymbols.otf
  install -vDm644 -t "$pkgdir/usr/share/fontconfig/conf.avail" font/10-powerline-symbols.conf
  install -vdm755 "$pkgdir/usr/share/fontconfig/conf.default"
  ln -vs /usr/share/fontconfig/conf.avail/10-powerline-symbols.conf \
    "$pkgdir/usr/share/fontconfig/conf.default/10-powerline-symbols.conf"
  install -vDm644 -t "$pkgdir/usr/share/licenses/powerline-fonts" LICENSE
}

package_vim-powerline() {
  pkgdesc="Vim plugin for powerline"
  groups=(vim-plugins)
  depends=(
    powerline
    vim-plugin-runtime
  )
  provides=(powerline-vim)
  replaces=(powerline-vim)

  cd $pkgbase
  install -vDm644 -t "$pkgdir/usr/share/vim/vimfiles/plugin" \
    powerline/bindings/vim/plugin/powerline.vim
  install -vDm644 -t "$pkgdir/usr/share/licenses/vim-powerline" LICENSE
}
