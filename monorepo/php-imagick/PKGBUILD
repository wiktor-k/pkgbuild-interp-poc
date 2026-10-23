# Maintainer: David Runge <dvzrv@archlinux.org>

_name=imagick
pkgbase=php-imagick
pkgname=(
  php-imagick
  php-legacy-imagick
)
pkgver=3.8.1
pkgrel=3
pkgdesc="PHP extension to create and modify images using the ImageMagick library"
arch=(x86_64)
url="https://github.com/imagick/imagick"
license=(
  0BSD
  PHP-3.01
)
depends=(
  libgomp
  glibc
  imagemagick
  ttf-font
)
makedepends=(
  librsvg
  php
  php-legacy
)
checkdepends=(ttf-dejavu)
optdepends=(
  'librsvg: for SVG support'
)
source=(
  $pkgbase-$pkgver.tar.gz::https://github.com/$_name/$_name/archive/refs/tags/$pkgver.tar.gz
  $_name.ini
  LICENSE
  $pkgbase-3.8.0-skip_version_check_by_default.patch
)
sha512sums=('b584394ce2c785a206211454208d23f5675e05162ccf76dacd267a1e4f6e8cb8d373dd3f660c42da5b399df4e89256be7d08465c3aa8d8666207fe61e0fadfd5'
            '720e050c7db81069b9f27624dd9430dd4952cfbc7f2374a58a6ce1eabd95bf936f3f9f9c53d44ff402acd38e869c2b9ec190b9d8a1e2aa204ecf9719e2708a42'
            'a33658d9271e5c537ccd41bf540b463ad2a5eca4a060c80486ff42a736f0aa042d10436e7177c34d792177cb11285243dee1f31c4df54fb0bfaabbc306406930'
            '3772be11c225f16830e386d030d6e92fb1fad86c4b4d29e66804d8c155f6b3c27a60a956379ba87faf0a746fe46c770726abd37f18927201dbb22b5a1d986e8c')
b2sums=('6d949f7198e0213d2812f8e1464576ea434e5bbd1044a9d3f96076e4cfc2779bcba48108d95b786b4de9e0168cb99424d9f133505f01d9376dd1cc69b2a7b342'
        '79f08a42378d0589f162cc091e7f5e216ded51063a9b0230b933c3eddf8427ebbe44095cbe1a26e2eb8cb12540edfae14b10298a2daef1ae73cb3133cd18e01b'
        'a29664104e1ee73ca0aee1d633e9095d92a57c92787f8d8740bdb7211ba3205782ed8677f539bdb8cae3dd75a3694be3132e185fa3fc4b3f401e1f88eb776101'
        'dfccdf3b2ef248b9617eb10b683ca09d4f663baa707d6ced8ede779c51c08ef0a26e76452367f1b8e19ad4835bf3b116bedeeaaf68c0b54a7983a6d217720edd')

prepare() {
  # disable the broken version check, which consistently provides false positives and leads to users flagging the package ood
  # we rely on imagemagick to communicate breaking changes with soname changes:
  # https://gitlab.archlinux.org/archlinux/packaging/packages/php-imagick/-/issues/3
  patch -Np1 -d $_name-$pkgver -i ../$pkgbase-3.8.0-skip_version_check_by_default.patch

  # setting package version: https://bugs.archlinux.org/task/64185
  sed -e "s/@PACKAGE_VERSION@/$pkgver/" -i $_name-$pkgver/php_imagick.h $_name-$pkgver/package.xml

  # remove broken tests
  local broken_tests=(
    "$_name-$pkgver/tests/013-read-filehandle-memory-stream.phpt"
    "$_name-$pkgver/tests/024-ispixelsimilar.phpt"
    "$_name-$pkgver/tests/151_Imagick_subImageMatch_basic.phpt"
    "$_name-$pkgver/tests/268_ImagickDraw_getDensity_basic.phpt"
    "$_name-$pkgver/tests/287_Imagick_GetImageChannelRange_basic.phpt"
    "$_name-$pkgver/tests/316_Imagick_getImageKurtosis.phpt"
    "$_name-$pkgver/tests/bug_73840.phpt"
  )
  printf 'Remove broken tests\n'
  rm -v "${broken_tests[@]}"

  cp -av $_name-$pkgver "${pkgname[1]}-$pkgver"
  (
    cd $_name-$pkgver
    phpize
  )
  (
    cd "${pkgname[1]}-$pkgver"
    phpize-legacy
  )
}

build() {
  (
    cd $_name-$pkgver
    ./configure --prefix=/usr
    make
  )
  (
    cd "${pkgname[1]}-$pkgver"
    ./configure --prefix=/usr
    make
  )
}

check() {
  export NO_INTERACTION=true
  make -k test -C $_name-$pkgver
  make -k test -C "${pkgname[1]}-$pkgver"
}

package_php-imagick() {
  depends+=(php)
  backup=(etc/php/conf.d/$_name.ini)

  cd $_name-$pkgver
  make INSTALL_ROOT="$pkgdir/" install
  install -vDm 644 ../$_name.ini -t "$pkgdir/etc/php/conf.d/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 {ChangeLog,CREDITS,README.md} -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 examples/*.php -t "$pkgdir/usr/share/doc/$pkgname/examples/"
}

package_php-legacy-imagick() {
  depends+=(php-legacy)
  backup=(etc/php-legacy/conf.d/$_name.ini)

  cd $pkgname-$pkgver
  make INSTALL_ROOT="$pkgdir/" install
  install -vDm 644 ../$_name.ini -t "$pkgdir/etc/php-legacy/conf.d/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
  install -vDm 644 {ChangeLog,CREDITS,README.md} -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 examples/*.php -t "$pkgdir/usr/share/doc/$pkgname/examples/"
}
