# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: XavierCLL <xavier.corredor.llano (a) gmail.com>
# Contributor: Christian Rebischke <chris.rebischke@archlinux.org>
# PKGBUILD base; Simper < simonpersson1 AT gmail DOT com >

pkgname=kup
pkgver=6.8.1
_dirver=$(echo $pkgver | cut -d. -f1-3)
pkgrel=1
pkgdesc='A backup software for helping people to keep up-to-date backups'
arch=(x86_64)
url='https://apps.kde.org/kup/'
license=(GPL-2.0-only
         GPL-2.0-or-later
         GPL-3.0-only)
depends=(glibc
         kcmutils
         kcompletion
         kconfig
         kconfigwidgets
         kcoreaddons
         kcrash
         kdbusaddons
         ki18n
         kidletime
         kio
         kirigami
         kjobwidgets
         knotifications
         kwidgetsaddons
         kwindowsystem
         kxmlgui
         libgcc
         libgit2
         libplasma
         libstdc++
         qt6-base
         qt6-declarative
         solid)
makedepends=(extra-cmake-modules
             git)
optdepends=('bup: bup backend'
            'rsync: support for synced backups')
groups=(plasma)
source=(https://download.kde.org/stable/plasma/$_dirver/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('867483361720b389fa9d4219e77fd51486157e1fd410587a86ab58b180607a77'
            'SKIP')
validpgpkeys=('E0A3EB202F8E57528E13E72FD7574483BB57B18D'  # Jonathan Esk-Riddell <jr@jriddell.org>
              '0AAC775BB6437A8D9AF7A3ACFE0784117FBCE11D'  # Bhushan Shah <bshah@kde.org>
              'D07BD8662C56CB291B316EB2F5675605C74E02CF'  # David Edmundson <davidedmundson@kde.org>
              '1FA881591C26B276D7A5518EEAAF29B42A678C20') # Marco Martin <notmart@gmail.com>

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DBUILD_TESTING=OFF \
    -DQT_MAJOR_VERSION=6
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
