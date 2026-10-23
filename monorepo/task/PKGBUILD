# Maintainer: Lukas Fleischer <lfleischer@archlinux.org>
# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Contributor: Peter Lewis <plewis@aur.archlinux.org>
# Contributor: TDY <tdy@gmx.com>
# Contributor: Ray Kohler <ataraxia@gmail.com>
# Contributor: muflax <muflax@gmail.com>
# Contributor: coolkehon <coolkehon@gmail.com>

pkgname=task
pkgver=3.5.0
_commit=3419d5ba1e6a780fbbb0a8d68d8fd31e675c1bf8
pkgrel=2
pkgdesc="Taskwarrior, a command-line todo list manager"
arch=('x86_64')
url="https://taskwarrior.org/"
license=('MIT')
depends=('util-linux' 'gnutls')
makedepends=('cmake' 'git' 'cargo')
optdepends=('bash-completion: for bash completion' 'python: for python export addon' 'ruby: for ruby export addon' 'perl: for perl export addon' 'perl-json: for perl export addon')
source=(
  "$pkgname::git+https://github.com/GothenburgBitFactory/taskwarrior.git#commit=$_commit"
  "$pkgname-libshared::git+https://github.com/GothenburgBitFactory/libshared.git"
)
sha256sums=('f857266a2be231365090fbc44e87536e93f79c33634ffb06182a679dc47fcc49'
            'SKIP')
options=('!lto')

prepare() {
  cd "$srcdir/$pkgname"
  git submodule init
  git config submodule."src/libshared".url "${srcdir}/${pkgname}"-libshared
  git -c protocol.file.allow=always submodule update --init --recursive
}

build() {
  cd "$srcdir/$pkgname"

  cmake -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_BUILD_TYPE=None .
  make
}

package() {
  cd "$srcdir/$pkgname"
  make DESTDIR="$pkgdir" install

  # Note that we rename the bash completion script for bash-completion > 1.99, until upstream does so.
  install -Dm644 "$pkgdir/usr/share/doc/task/scripts/bash/task.sh" "$pkgdir/usr/share/bash-completion/completions/task"

  install -Dm644 "$pkgdir/usr/share/doc/task/scripts/vim/ftdetect/task.vim" "$pkgdir/usr/share/vim/vimfiles/ftdetect/task.vim"
  install -Dm644 "$pkgdir/usr/share/doc/task/scripts/vim/syntax/taskdata.vim" "$pkgdir/usr/share/vim/vimfiles/syntax/taskdata.vim"
  install -Dm644 "$pkgdir/usr/share/doc/task/scripts/vim/syntax/taskedit.vim" "$pkgdir/usr/share/vim/vimfiles/syntax/taskedit.vim"
  install -Dm644 "$pkgdir/usr/share/doc/task/scripts/vim/syntax/taskrc.vim" "$pkgdir/usr/share/vim/vimfiles/syntax/taskrc.vim"

  install -Dm644 "$srcdir/$pkgname/LICENSE" "$pkgdir/usr/share/licenses/task/LICENSE"
}
